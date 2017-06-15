package pig

import (
	"github.com/jkomoros/boardgame"
)

//+autoreader
type gameState struct {
	CurrentPlayer boardgame.PlayerIndex
	Die           *boardgame.SizedStack
}

//+autoreader
type playerState struct {
	playerIndex boardgame.PlayerIndex
	RoundScore  int
	TotalScore  int
}

func concreteStates(state boardgame.State) (*gameState, []*playerState) {
	game := state.Game().(*gameState)

	players := make([]*playerState, len(state.Players()))

	for i, player := range state.Players() {
		players[i] = player.(*playerState)
	}

	return game, players
}

func (p *playerState) PlayerIndex() boardgame.PlayerIndex {
	return p.playerIndex
}
