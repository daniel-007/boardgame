package playingcards

import (
	"github.com/jkomoros/boardgame"
	"testing"
)

func TestNewDeck(t *testing.T) {

	chest := boardgame.NewComponentChest()

	deck := NewDeck(false)

	chest.AddDeck("cards", deck)

	if len(deck.Components()) != 52 {
		t.Error("We asked for no jokers but got wrong number of cards", len(deck.Components()))
	}

	checkExpectedRun(deck, 0, t)

	withJokers := NewDeck(true)

	chest.AddDeck("jokers", withJokers)

	if len(withJokers.Components()) != 54 {
		t.Error("Deck with jokers had wrong number of cards:", len(withJokers.Components()))
	}

}

//Checks that the deck, at starting Index, has the 52 main cards in canonical order.
func checkExpectedRun(deck *boardgame.Deck, startingIndex int, t *testing.T) {

	if len(deck.Components()) < 52+startingIndex {
		t.Error("Deck didn't have enough items")
	}

	suits := []Suit{SuitSpades, SuitHearts, SuitClubs, SuitDiamonds}

	expectedRank := RankAce
	expectedSuitIndex := 0
	expectedSuit := suits[expectedSuitIndex]

	components := deck.Components()

	for i := startingIndex; i < (startingIndex + 52); i++ {
		card := components[i].Values.(*Card)

		if card.Rank != expectedRank {
			t.Error("Card", i, "had wrong rank. Wanted", expectedRank, "Got", card.Rank)
		}

		if card.Suit != expectedSuit {
			t.Error("Card", i, "had wrong suit. Wanted", expectedSuit, "Got", card.Suit)
		}

		expectedRank++
		if expectedRank > RankKing {
			expectedRank = RankAce
			expectedSuitIndex++
			if expectedSuitIndex < len(suits) {
				expectedSuit = suits[expectedSuitIndex]
			}
		}
	}

}
