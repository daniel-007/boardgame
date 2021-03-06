/*

	extendedgame is the definition of a StorageRecord for ExtendedGame. In a
	separate package to avoid dependency cycles.

*/
package extendedgame

import (
	"encoding/json"
	"github.com/jkomoros/boardgame"
)

type StorageRecord struct {
	Open    bool
	Visible bool
	Owner   string
}

type CombinedStorageRecord struct {
	boardgame.GameStorageRecord
	StorageRecord
}

func DefaultStorageRecord() *StorageRecord {
	return &StorageRecord{
		Open:    true,
		Visible: true,
		Owner:   "",
	}
}

func (c *CombinedStorageRecord) String() string {
	blob, _ := json.Marshal(c)
	return string(blob) + "\n"
}
