package checkers

import (
	"errors"
	"github.com/jkomoros/boardgame"
	"github.com/jkomoros/boardgame/moves"
)

//+autoreader
type MovePlaceToken struct {
	moves.FixUpMulti
	TargetIndex int
}

func (m *MovePlaceToken) DefaultsForState(state boardgame.State) {

	game := state.GameState().(*gameState)

	if game.UnusedTokens.NumComponents() <= 0 {
		return
	}

	nextToken := game.UnusedTokens.ComponentAt(0)

	nextTokenVals := nextToken.Values.(*token)

	//Red starts at top
	fromBottom := false

	if nextTokenVals.Color.Value() == ColorBlack {
		fromBottom = true
	}

	startIndex := 0
	increment := 1
	endCondition := game.Spaces.Len()

	if fromBottom {
		startIndex = game.Spaces.Len() - 1
		increment = -1
		endCondition = 0
	}

	for i := startIndex; i != endCondition; i += increment {
		//We're only allowed to put tokens on black spaces
		if !spaceIsBlack(i) {
			continue
		}
		if game.Spaces.ComponentAt(i) == nil {
			m.TargetIndex = i
			return
		}
	}

}

func (m *MovePlaceToken) Legal(state boardgame.State, proposer boardgame.PlayerIndex) error {
	if err := m.FixUpMulti.Legal(state, proposer); err != nil {
		return err
	}

	game := state.GameState().(*gameState)

	if game.UnusedTokens.NumComponents() == 0 {
		return errors.New("No more components to place")
	}

	if game.Spaces.ComponentAt(m.TargetIndex) != nil {
		return errors.New("That space is already filled")
	}

	if !spaceIsBlack(m.TargetIndex) {
		return errors.New("The proposed space is not black")
	}

	return nil
}

func (m *MovePlaceToken) Apply(state boardgame.MutableState) error {
	game := state.GameState().(*gameState)
	return game.UnusedTokens.MoveComponent(boardgame.FirstComponentIndex, game.Spaces, m.TargetIndex)
}

type MoveMoveToken struct {
	moves.CurrentPlayer
	TokenIndexToMove int
	SpaceIndex       int
}

func (m *MoveMoveToken) Legal(state boardgame.State, proposer boardgame.PlayerIndex) error {
	if err := m.CurrentPlayer.Legal(state, proposer); err != nil {
		return err
	}

	p := state.CurrentPlayer().(*playerState)

	g := state.GameState().(*gameState)

	c := g.Spaces.ComponentAt(m.TokenIndexToMove)

	if c == nil {
		return errors.New("That space does not have a component in it")
	}

	t := c.Values.(*token)

	if !p.Color.Equals(t.Color) {
		return errors.New("That token isn't your token to move!")
	}

	if !spaceIsBlack(m.SpaceIndex) {
		return errors.New("You can only move to spaces that are black.")
	}

	if g.Spaces.ComponentAt(m.SpaceIndex) != nil {
		return errors.New("The space you're trying to move to is occupied.")
	}

	//TODO: make sure the move is legal via graph connectedness and direction
	//(depending on Crowned).

	return nil

}

func (m *MoveMoveToken) Apply(state boardgame.MutableState) error {

	g := state.GameState().(*gameState)

	if err := g.Spaces.SwapComponents(m.TokenIndexToMove, m.SpaceIndex); err != nil {
		return errors.New("Couldn't move token: " + err.Error())
	}

	//TODO: capture any tokens moved over

	//TODO: mark turn as done, if we didn't capture a token.

	return nil

}