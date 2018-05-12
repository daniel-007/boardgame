package moves

import (
	"github.com/jkomoros/boardgame"
	"github.com/jkomoros/boardgame/moves/auto"
	"github.com/workfit/tester/assert"
	"testing"
)

//+autoreader
type moveDealCards struct {
	DealCountComponents
}

func (m *moveDealCards) TargetCount() int {
	return 2
}

func (m *moveDealCards) GameStack(gState boardgame.MutableSubState) boardgame.MutableStack {
	return gState.(*gameState).DrawStack
}

func (m *moveDealCards) PlayerStack(pState boardgame.MutablePlayerState) boardgame.MutableStack {
	return pState.(*playerState).Hand
}

//+autoreader
type moveDealOtherCards struct {
	DealCountComponents
}

func (m *moveDealOtherCards) TargetCount() int {
	return 3
}

func (m *moveDealOtherCards) GameStack(gState boardgame.MutableSubState) boardgame.MutableStack {
	return gState.(*gameState).DrawStack
}

func (m *moveDealOtherCards) PlayerStack(pState boardgame.MutablePlayerState) boardgame.MutableStack {
	return pState.(*playerState).OtherHand
}

//+autoreader
type moveCurrentPlayerDraw struct {
	CurrentPlayer
}

func (m *moveCurrentPlayerDraw) Apply(state boardgame.MutableState) error {
	game, players := concreteStates(state)

	p := players[game.CurrentPlayer]

	return game.DrawStack.MutableFirst().MoveTo(p.Hand, boardgame.FirstSlotIndex)
}

//+autoreader
type moveStartPhaseDrawAgain struct {
	StartPhase
}

func (m *moveStartPhaseDrawAgain) PhaseToStart(currentPhase int) int {
	return phaseDrawAgain
}

func defaultMoveInstaller(manager *boardgame.GameManager) *boardgame.MoveTypeConfigBundle {

	return boardgame.NewMoveTypeConfigBundle().AddOrderedMovesForPhase(phaseSetUp,
		auto.MustConfig(
			new(moveDealCards),
			WithMoveName("Deal Components From Game Stack DrawStack To Player Stack Hand To Each Player 2 Times"),
		),
		auto.MustConfig(
			new(moveDealOtherCards),
			WithMoveName("Deal Other Cards OVERRIDE"),
		),
		auto.MustConfig(
			new(StartPhase),
			WithPhaseToStart(phaseNormalPlay, phaseEnum),
		),
	).AddMovesForPhase(phaseNormalPlay,
		auto.MustConfig(
			new(moveCurrentPlayerDraw),
			WithMoveName("Draw Card"),
		),
		auto.MustConfig(
			new(moveStartPhaseDrawAgain),
			WithMoveName("Start Phase Draw Again"),
			WithIsFixUp(false),
		),
	).AddOrderedMovesForPhase(phaseDrawAgain,
		auto.MustConfig(
			new(DealComponentsUntilPlayerCountReached),
			WithMoveName("Deal Cards To Three"),
			WithGameStack("DrawStack"),
			WithPlayerStack("Hand"),
			WithTargetCount(3),
		),
	)
}

func TestGeneral(t *testing.T) {
	manager, err := newGameManager(defaultMoveInstaller)

	assert.For(t).ThatActual(err).IsNil()

	game := manager.NewGame()

	err = game.SetUp(0, nil, nil)

	assert.For(t).ThatActual(err).IsNil()

	//4 players, 2 rounds for inital cards, then 4 * 3 for other cards, then
	//NewStartPhase.
	assert.For(t).ThatActual(game.Version()).Equals(21)

	gameState, playerStates := concreteStates(game.CurrentState())

	assert.For(t).ThatActual(gameState.DrawStack.NumComponents()).Equals(52 - 20)
	assert.For(t).ThatActual(gameState.Phase.Value()).Equals(phaseNormalPlay)

	for i, player := range playerStates {
		assert.For(t, i).ThatActual(player.Hand.NumComponents()).Equals(2)
		assert.For(t, i).ThatActual(player.OtherHand.NumComponents()).Equals(3)
	}

	historicalMovesCount(t,
		[]string{
			"Deal Components From Game Stack DrawStack To Player Stack Hand To Each Player 2 Times",
			"Deal Other Cards OVERRIDE",
			"Start Phase Normal Play",
		},
		[]int{
			8,
			12,
			1,
		}, game.MoveRecords(-1))

	move := game.MoveByName("Draw Card")

	assert.For(t).ThatActual(move).IsNotNil()

	err = <-game.ProposeMove(move, 0)

	assert.For(t).ThatActual(err).IsNil()

	move = game.MoveByName("Start Phase Draw Again")

	assert.For(t).ThatActual(move).IsNotNil()

	err = <-game.ProposeMove(move, 0)

	assert.For(t).ThatActual(err).IsNil()

	assert.For(t).ThatActual(manager.Delegate().CurrentPhase(game.CurrentState())).Equals(phaseDrawAgain)

	//3 additional moves, but skipping the one player who already had 3 in their hand.
	assert.For(t).ThatActual(game.Version()).Equals(26)

}

func historicalMovesCount(t *testing.T, moveNames []string, counts []int, records []*boardgame.MoveStorageRecord) {
	if len(moveNames) != len(counts) {
		t.Error("MoveNames and counts did not match length")
	}

	currentMoveIndex := 0
	counterInMoveType := 0

	for i, move := range records {

		if counterInMoveType >= counts[currentMoveIndex] {
			currentMoveIndex++
			counterInMoveType = 0
			if currentMoveIndex > len(moveNames) {
				t.Error("Fell off end of configuration")
			}
		}

		if move.Name != moveNames[currentMoveIndex] {
			t.Error("Unexpected move at ", i, move.Name, moveNames[currentMoveIndex])
		}

		counterInMoveType++
	}
}
