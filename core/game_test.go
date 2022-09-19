package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentPlayerCards(t *testing.T) {
	t.Run("Returns current player's cards", func(t *testing.T) {
		game := newGame()

		playerCards := game.CurrentPlayerCards()

		wantedCards := game.currentPlayer.cards.Copy()
		wantedCards[GoodCamel] = game.currentPlayer.herdSize

		assert.Equal(t, wantedCards, playerCards)
	})

	t.Run("Mutating the result won't affect game", func(t *testing.T) {
		game := newGame()

		playerCards := game.CurrentPlayerCards()
		playerCards[GoodDiamond] = 50

		assert.NotEqual(t, game.currentPlayer.cards, playerCards)
	})
}

func TestCardsOnTable(t *testing.T) {
	t.Run("Returns cards on table", func(t *testing.T) {
		game := newGame()

		cards := game.CardsOnTable()

		assert.Equal(t, game.cardsOnTable, cards)
	})

	t.Run("Mutating the result won't affect game", func(t *testing.T) {
		game := newGame()

		cards := game.CardsOnTable()
		cards[GoodDiamond] = 50

		assert.NotEqual(t, game.cardsOnTable, cards)
	})
}

func TestCurrentPlayerName(t *testing.T) {
	game := newGame()

	name := game.CurrentPlayerName()

	assert.Equal(t, game.currentPlayer.name, name)
}
