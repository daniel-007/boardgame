/*

bolt provides a bolt-backed database that implements both
boardgame.StorageManager and boardgame/server.StorageManager.

*/
package bolt

import (
	"encoding/json"
	"errors"
	"github.com/boltdb/bolt"
	"github.com/jkomoros/boardgame"
	"github.com/jkomoros/boardgame/server/api/users"
	"log"
	"os"
	"strconv"
	"strings"
)

//TODO: test this package

type StorageManager struct {
	db       *bolt.DB
	filename string
}

var (
	statesBucket      = []byte("States")
	gamesBucket       = []byte("Games")
	usersBucket       = []byte("Users")
	cookiesBucket     = []byte("Cookies")
	gameUsersBucket   = []byte("GameUsers")
	agentStatesBucket = []byte("AgentStates")
)

func NewStorageManager(fileName string) *StorageManager {
	db, err := bolt.Open(fileName, 0600, nil)

	if err != nil {
		panic("Couldn't open db")
	}

	err = db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists(gamesBucket); err != nil {
			return errors.New("Cannot create games bucket" + err.Error())
		}
		if _, err := tx.CreateBucketIfNotExists(statesBucket); err != nil {
			return errors.New("Cannot create states bucket" + err.Error())
		}
		if _, err := tx.CreateBucketIfNotExists(usersBucket); err != nil {
			return errors.New("Cannot create users bucket" + err.Error())
		}
		if _, err := tx.CreateBucketIfNotExists(cookiesBucket); err != nil {
			return errors.New("Cannot create cookies bucket" + err.Error())
		}
		if _, err := tx.CreateBucketIfNotExists(gameUsersBucket); err != nil {
			return errors.New("Cannot create game users bucket" + err.Error())
		}
		if _, err := tx.CreateBucketIfNotExists(agentStatesBucket); err != nil {
			return errors.New("Cannot create agent states bucket" + err.Error())
		}
		return nil
	})

	if err != nil {
		//Not able to initalize DB
		return nil
	}
	//We don't defer DB close; our users need to.
	return &StorageManager{
		db:       db,
		filename: fileName,
	}

}

func keyForState(gameId string, version int) []byte {
	return []byte(gameId + "_" + strconv.Itoa(version))
}

func keyForGame(id string) []byte {
	return []byte(strings.ToUpper(id))
}

func keyForUser(uid string) []byte {
	return []byte(uid)
}

func keyForCookie(cookie string) []byte {
	return []byte(cookie)
}

func keyForAgentState(gameId string, player boardgame.PlayerIndex) []byte {
	return []byte(gameId + "-" + player.String())
}

func (s *StorageManager) Name() string {
	return "bolt"
}

func (s *StorageManager) State(gameId string, version int) (boardgame.StateStorageRecord, error) {
	if gameId == "" {
		return nil, errors.New("No game provided")
	}

	if version < 0 {
		return nil, errors.New("Invalid version")
	}

	var record []byte

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(statesBucket)

		if b == nil {
			return errors.New("Couldn't get bucket")
		}

		record = b.Get(keyForState(gameId, version))
		return nil
	})

	if err != nil {
		return nil, err
	}

	if record == nil {
		return nil, errors.New("No such version for game")
	}

	return record, nil
}

func (s *StorageManager) Game(id string) (*boardgame.GameStorageRecord, error) {

	var rawRecord []byte

	err := s.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(gamesBucket)
		if b == nil {
			return errors.New("Couldn't open bucket")
		}
		rawRecord = b.Get(keyForGame(id))
		return nil
	})

	if err != nil {
		return nil, errors.New("Transacation error " + err.Error())
	}

	if rawRecord == nil {
		return nil, errors.New("No such game found")
	}

	var record boardgame.GameStorageRecord

	if err := json.Unmarshal(rawRecord, &record); err != nil {
		return nil, errors.New("Unmarshal error " + err.Error())
	}

	return &record, nil

}

func (s *StorageManager) SaveGameAndCurrentState(game *boardgame.GameStorageRecord, state boardgame.StateStorageRecord) error {

	version := game.Version

	serializedGameRecord, err := json.Marshal(game)

	if err != nil {
		return errors.New("Couldn't serialize the internal game record: " + err.Error())
	}

	return s.db.Update(func(tx *bolt.Tx) error {
		gBucket := tx.Bucket(gamesBucket)

		if gBucket == nil {
			return errors.New("Couldn't open games bucket")
		}

		sBucket := tx.Bucket(statesBucket)

		if sBucket == nil {
			return errors.New("Could open states bucket")
		}

		if err := gBucket.Put(keyForGame(game.Id), serializedGameRecord); err != nil {
			return err
		}

		if err := sBucket.Put(keyForState(game.Id, version), state); err != nil {
			return err
		}

		return nil

	})

}

func (s *StorageManager) AgentState(gameId string, player boardgame.PlayerIndex) ([]byte, error) {

	var result []byte

	err := s.db.View(func(tx *bolt.Tx) error {

		aBucket := tx.Bucket(agentStatesBucket)

		if aBucket == nil {
			return errors.New("Couldn't open agent states bucket")
		}

		result = aBucket.Get(keyForAgentState(gameId, player))
		return nil

	})

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (s *StorageManager) SaveAgentState(gameId string, player boardgame.PlayerIndex, state []byte) error {

	return s.db.Update(func(tx *bolt.Tx) error {
		aBucket := tx.Bucket(agentStatesBucket)

		if aBucket == nil {
			return errors.New("Couldn't open agent states bucket")
		}

		if err := aBucket.Put(keyForAgentState(gameId, player), state); err != nil {
			return err
		}
		return nil
	})

}

func (s *StorageManager) ListGames(max int) []*boardgame.GameStorageRecord {

	var result []*boardgame.GameStorageRecord

	err := s.db.View(func(tx *bolt.Tx) error {

		gBucket := tx.Bucket(gamesBucket)

		if gBucket == nil {
			return errors.New("couldn't open games bucket")
		}

		c := gBucket.Cursor()
		for k, v := c.First(); k != nil; k, v = c.Next() {
			if len(result) >= max {
				break
			}

			var record boardgame.GameStorageRecord

			if err := json.Unmarshal(v, &record); err != nil {
				return errors.New("Couldn't deserialize a game: " + err.Error())
			}

			result = append(result, &record)
		}

		return nil

	})

	if err != nil {
		return nil
	}

	return result

}

func (s *StorageManager) SetPlayerForGame(gameId string, playerIndex boardgame.PlayerIndex, userId string) error {

	ids := s.UserIdsForGame(gameId)

	if ids == nil {
		return errors.New("Couldn't fetch original player indexes for that game")
	}

	if int(playerIndex) < 0 || int(playerIndex) >= len(ids) {
		return errors.New("PlayerIndex " + playerIndex.String() + " is not valid for this game")
	}

	if ids[playerIndex] != "" {
		return errors.New("PlayerIndex " + playerIndex.String() + " is already taken")
	}

	user := s.GetUserById(userId)

	if user == nil {
		return errors.New("That userId does not describe an existing user")
	}

	ids[playerIndex] = userId

	err := s.db.Update(func(tx *bolt.Tx) error {
		gUBucket := tx.Bucket(gameUsersBucket)

		if gUBucket == nil {
			return errors.New("Couldn't open game useres bucket")
		}

		blob, err := json.Marshal(ids)

		if err != nil {
			return errors.New("Unable to marshal ids blob: " + err.Error())
		}

		return gUBucket.Put(keyForGame(gameId), blob)
	})

	if err != nil {
		return errors.New("Unable to form association: " + err.Error())
	}

	return nil

}

func (s *StorageManager) UserIdsForGame(gameId string) []string {

	noRecordErr := errors.New("No such record")

	var result []string

	err := s.db.View(func(tx *bolt.Tx) error {
		gUBucket := tx.Bucket(gameUsersBucket)

		if gUBucket == nil {
			return errors.New("Couldn't open game users bucket")
		}

		blob := gUBucket.Get(keyForGame(gameId))

		if blob == nil {
			//NO such game info.
			return noRecordErr
		}

		return json.Unmarshal(blob, &result)
	})

	if err == noRecordErr {
		//It's possible that we just haven't stored anything for this user before.

		gameRecord, err := s.Game(gameId)

		if err != nil {
			log.Println("Couldn fetch game: " + err.Error())
			return nil
		}

		if gameRecord == nil {
			log.Println("No such game")
			return nil
		}

		return make([]string, gameRecord.NumPlayers)
	}

	if err != nil {
		log.Println("Error in UserIdsForGame: ", err)
		return nil
	}

	return result
}

func (s *StorageManager) UpdateUser(user *users.StorageRecord) error {
	err := s.db.Update(func(tx *bolt.Tx) error {

		uBucket := tx.Bucket(usersBucket)

		if uBucket == nil {
			return errors.New("couldn't open users bucket")
		}

		blob, err := json.Marshal(user)

		if err != nil {
			return errors.New("Couldn't marshal user: " + err.Error())
		}

		return uBucket.Put(keyForUser(user.Id), blob)

	})

	return err
}

func (s *StorageManager) GetUserById(uid string) *users.StorageRecord {
	var result users.StorageRecord

	err := s.db.View(func(tx *bolt.Tx) error {
		uBucket := tx.Bucket(usersBucket)

		if uBucket == nil {
			return errors.New("Couldn't open users bucket")
		}

		uBlob := uBucket.Get(keyForUser(uid))

		if uBlob == nil {
			return errors.New("No such user")
		}

		return json.Unmarshal(uBlob, &result)
	})

	if err != nil {
		log.Println("Failure in GetUserById: ", err)
		return nil
	}

	return &result
}

func (s *StorageManager) GetUserByCookie(cookie string) *users.StorageRecord {

	var result users.StorageRecord

	err := s.db.View(func(tx *bolt.Tx) error {

		cBucket := tx.Bucket(cookiesBucket)

		if cBucket == nil {
			return errors.New("Couldn't open cookies bucket")
		}

		c := cBucket.Get(keyForCookie(cookie))

		if c == nil {
			return errors.New("No such cookie: " + cookie)
		}

		uBucket := tx.Bucket(usersBucket)

		if uBucket == nil {
			return errors.New("couldn't open users bucket")
		}

		uBlob := uBucket.Get(keyForUser(string(c)))

		if uBlob == nil {
			return errors.New("The user specified by cookie was not found")
		}

		if err := json.Unmarshal(uBlob, &result); err != nil {
			return errors.New("Unable to unmarshal user objet: " + err.Error())
		}

		return nil

	})

	if err != nil {
		log.Println("Failure in GetUserByCookie", err)
		return nil
	}

	return &result

}

func (s *StorageManager) ConnectCookieToUser(cookie string, user *users.StorageRecord) error {
	err := s.db.Update(func(tx *bolt.Tx) error {

		cBucket := tx.Bucket(cookiesBucket)

		if cBucket == nil {
			return errors.New("couldn't open cookies bucket")
		}

		if user == nil {
			//Delete the cookie.
			return cBucket.Delete(keyForCookie(cookie))
		}

		return cBucket.Put(keyForCookie(cookie), keyForUser(user.Id))

	})

	return err
}

func (s *StorageManager) Connect(config string) error {
	return nil
}

func (s *StorageManager) Close() {
	s.db.Close()
}

func (s *StorageManager) CleanUp() {
	os.Remove(s.filename)
}
