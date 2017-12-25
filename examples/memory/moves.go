package memory

import (
	"errors"
	"github.com/jkomoros/boardgame"
	"github.com/jkomoros/boardgame/moves"
	"time"
)

const HideCardsDuration = 4 * time.Second

/**************************************************
 *
 * MoveRevealCard Implementation
 *
 **************************************************/

//+autoreader
type MoveRevealCard struct {
	moves.CurrentPlayer
	CardIndex int
}

func (m *MoveRevealCard) DefaultsForState(state boardgame.State) {

	m.CurrentPlayer.DefaultsForState(state)

	game, _ := concreteStates(state)

	for i, c := range game.HiddenCards.Components() {
		if c != nil {
			m.CardIndex = i
			break
		}
	}
}

func (m *MoveRevealCard) Legal(state boardgame.State, proposer boardgame.PlayerIndex) error {

	if err := m.CurrentPlayer.Legal(state, proposer); err != nil {
		return err
	}

	game, players := concreteStates(state)

	p := players[game.CurrentPlayer]

	if p.CardsLeftToReveal < 1 {
		return errors.New("You have no cards left to reveal this turn")
	}

	if m.CardIndex < 0 || m.CardIndex >= game.HiddenCards.Len() {
		return errors.New("Illegal card index.")
	}

	if game.HiddenCards.ComponentAt(m.CardIndex) == nil {
		if game.VisibleCards.ComponentAt(m.CardIndex) == nil {
			return errors.New("There is no card at that index.")
		} else {
			return errors.New("That card has already been revealed.")
		}
	}

	return nil
}

func (m *MoveRevealCard) Apply(state boardgame.MutableState) error {
	game, players := concreteStates(state)

	p := players[game.CurrentPlayer]

	p.CardsLeftToReveal--
	game.HiddenCards.MoveComponent(m.CardIndex, game.VisibleCards, m.CardIndex)

	//If the cards are the same, the FixUpMove CaptureCards will fire after this.

	return nil
}

/**************************************************
 *
 * MoveStartHideCardsTimer Implementation
 *
 **************************************************/

//+autoreader
type MoveStartHideCardsTimer struct {
	moves.FixUp
}

func (m *MoveStartHideCardsTimer) Legal(state boardgame.State, proposer boardgame.PlayerIndex) error {

	if err := m.Base.Legal(state, proposer); err != nil {
		return err
	}

	game, _ := concreteStates(state)

	if game.VisibleCards.NumComponents() != 2 {
		return errors.New("There aren't two cards showing!")
	}

	if game.HideCardsTimer.Active() {
		return errors.New("The timer is already active.")
	}

	var revealedCards []*boardgame.Component

	for _, c := range game.VisibleCards.Components() {
		if c != nil {
			revealedCards = append(revealedCards, c)
		}
	}

	cardOneType := revealedCards[0].Values.(*cardValue).Type
	cardTwoType := revealedCards[1].Values.(*cardValue).Type

	if cardOneType == cardTwoType {
		return errors.New("The two revealed cards are of the same type")
	}

	return nil
}

func (m *MoveStartHideCardsTimer) Apply(state boardgame.MutableState) error {
	game, _ := concreteStates(state)

	moveType := state.Game().Manager().MoveTypeByName(moveHideCardsConfig.Name)

	move := moveType.NewMove(state)

	game.HideCardsTimer.Start(HideCardsDuration, move)

	return nil
}

/**************************************************
 *
 * MoveCaptureCards Implementation
 *
 **************************************************/

//+autoreader
type MoveCaptureCards struct {
	moves.FixUp
}

func (m *MoveCaptureCards) Legal(state boardgame.State, proposer boardgame.PlayerIndex) error {
	if err := m.Base.Legal(state, proposer); err != nil {
		return err
	}

	game, _ := concreteStates(state)

	if game.VisibleCards.NumComponents() != 2 {
		return errors.New("There aren't two cards showing!")
	}

	var revealedCards []*boardgame.Component

	for _, c := range game.VisibleCards.Components() {
		if c != nil {
			revealedCards = append(revealedCards, c)
		}
	}

	cardOneType := revealedCards[0].Values.(*cardValue).Type
	cardTwoType := revealedCards[1].Values.(*cardValue).Type

	if cardOneType != cardTwoType {
		return errors.New("The two revealed cards are not of the same type")
	}

	return nil
}

func (m *MoveCaptureCards) Apply(state boardgame.MutableState) error {
	game, players := concreteStates(state)

	p := players[game.CurrentPlayer]

	for i, c := range game.VisibleCards.Components() {
		if c != nil {
			game.VisibleCards.MoveComponent(i, p.WonCards, boardgame.NextSlotIndex)
		}
	}

	return nil
}

/**************************************************
 *
 * MoveHideCards Implementation
 *
 **************************************************/

//+autoreader
type MoveHideCards struct {
	moves.CurrentPlayer
}

var moveHideCardsConfig = boardgame.MoveTypeConfig{
	Name:     "Hide Cards",
	HelpText: "After the current player has revealed both cards and tried to memorize them, this move hides the cards so that play can continue to next player.",
	MoveConstructor: func() boardgame.Move {
		return new(MoveHideCards)
	},
}

func (m *MoveHideCards) Legal(state boardgame.State, proposer boardgame.PlayerIndex) error {

	if err := m.CurrentPlayer.Legal(state, proposer); err != nil {
		return err
	}

	game, players := concreteStates(state)

	p := players[game.CurrentPlayer]

	if p.CardsLeftToReveal > 0 {
		return errors.New("You still have to reveal more cards before your turn is over")
	}

	if game.VisibleCards.NumComponents() < 1 {
		return errors.New("No cards left to hide!")
	}

	return nil
}

func (m *MoveHideCards) Apply(state boardgame.MutableState) error {
	game, _ := concreteStates(state)

	//Cancel a timer in case it was still going.
	game.HideCardsTimer.Cancel()

	for i, c := range game.VisibleCards.Components() {
		if c != nil {
			if err := game.VisibleCards.MoveComponent(i, game.HiddenCards, i); err != nil {
				return errors.New("Couldn't move component: " + err.Error())
			}
		}
	}

	return nil
}
