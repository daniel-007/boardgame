package memory

import (
	"errors"
	"github.com/jkomoros/boardgame"
	"github.com/jkomoros/boardgame/moves"
)

func init() {

	//Make sure that we get compile-time errors if our player and game state
	//don't adhere to the interfaces that moves.FinishTurn expects
	moves.VerifyFinishTurnStates(&gameState{}, &playerState{})
}

//+autoreader
type gameState struct {
	state          boardgame.State
	CurrentPlayer  boardgame.PlayerIndex
	HiddenCards    *boardgame.SizedStack `sanitize:"order"`
	RevealedCards  *boardgame.SizedStack
	HideCardsTimer *boardgame.Timer
}

//+autoreader
type playerState struct {
	state             boardgame.State
	playerIndex       boardgame.PlayerIndex
	CardsLeftToReveal int
	WonCards          *boardgame.GrowableStack `stack:"cards"`
}

func concreteStates(state boardgame.State) (*gameState, []*playerState) {
	game := state.GameState().(*gameState)

	players := make([]*playerState, len(state.PlayerStates()))

	for i, player := range state.PlayerStates() {
		players[i] = player.(*playerState)
	}

	return game, players
}

func (g *gameState) SetState(state boardgame.State) {
	g.state = state
}

func (p *playerState) SetState(state boardgame.State) {
	p.state = state
}

func (p *playerState) PlayerIndex() boardgame.PlayerIndex {
	return p.playerIndex
}

func (p *playerState) TurnDone() error {
	if p.CardsLeftToReveal > 0 {
		return errors.New("they still have cards left to reveal")
	}

	game, _ := concreteStates(p.state)

	if game.RevealedCards.NumComponents() > 0 {
		return errors.New("there are still some cards revealed, which they must hide")
	}

	return nil
}

func (p *playerState) ResetForTurnStart() error {
	p.CardsLeftToReveal = 2
	return nil
}

func (p *playerState) ResetForTurnEnd() error {
	return nil
}

func (g *gameState) CardsInGrid() int {
	return g.HiddenCards.NumComponents() + g.RevealedCards.NumComponents()
}

func (g *gameState) SetCurrentPlayer(currentPlayer boardgame.PlayerIndex) {
	g.CurrentPlayer = currentPlayer
}
