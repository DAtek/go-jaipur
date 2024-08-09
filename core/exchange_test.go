package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	t.Run("Player has the wanted cards", func(t *testing.T) {
		game := newGameMock()
		game.currentPlayer.cards = GoodMap{
			GoodCloth:   4,
			GoodLeather: 3,
		}
		game.currentPlayer.herdSize = 1
		game.cardsOnTable = GoodMap{GoodDiamond: 2, GoodGold: 2, GoodCamel: 1}

		error := game.Exchange(
			GoodMap{GoodDiamond: 2, GoodGold: 2, GoodCamel: 1},
			GoodMap{GoodCloth: 2, GoodLeather: 3},
		)

		assert.Nil(t, error)
		assert.Equal(t, Amount(2), game.player1.cards[GoodDiamond])
		assert.Equal(t, Amount(2), game.player1.cards[GoodGold])
		assert.Equal(t, Amount(2), game.player1.herdSize)
	})

	t.Run("Player doesn't have the sold cards", func(t *testing.T) {
		game := newGameMock()
		game.player1.herdSize = 1

		game.Exchange(
			GoodMap{GoodDiamond: 2, GoodGold: 1},
			GoodMap{GoodCloth: 1, GoodLeather: 1, GoodCamel: 1},
		)

		assert.Equal(t, Amount(1), game.player1.cards[GoodCloth])
		assert.Equal(t, Amount(0), game.player1.cards[GoodLeather])
		assert.Equal(t, Amount(0), game.player1.herdSize)
	})

	t.Run("The picked cards aren't on the table", func(t *testing.T) {
		game := newGameMock()

		game.Exchange(
			GoodMap{GoodDiamond: 2, GoodGold: 1},
			GoodMap{GoodCloth: 2, GoodLeather: 1},
		)

		assert.Equal(t, Amount(0), game.cardsOnTable[GoodDiamond])
		assert.Equal(t, Amount(0), game.cardsOnTable[GoodGold])
	})

	t.Run("Table has the dropped cards", func(t *testing.T) {
		game := newGameMock()

		game.Exchange(
			GoodMap{GoodDiamond: 2, GoodGold: 1},
			GoodMap{GoodCloth: 2, GoodLeather: 1},
		)

		assert.Equal(t, Amount(2), game.cardsOnTable[GoodCloth])
		assert.Equal(t, Amount(1), game.cardsOnTable[GoodLeather])
	})

	t.Run("Current player changes", func(t *testing.T) {
		game := newGameMock()

		game.Exchange(
			GoodMap{GoodDiamond: 2, GoodGold: 1},
			GoodMap{GoodCloth: 2, GoodLeather: 1},
		)

		assert.Equal(t, game.player2.name, game.currentPlayer.name)
	})

	t.Run("Error if player doesn't have proper cards", func(t *testing.T) {
		game := newGameMock()
		game.player1.cards = GoodMap{}
		game.cardsOnTable = GoodMap{GoodCloth: 2}

		error := game.Exchange(GoodMap{GoodCloth: 2}, GoodMap{GoodDiamond: 2})

		assert.EqualError(t, error, NotEnoughCardsToSellError.Error())
	})

	t.Run("Error if player doesn't big enough herd", func(t *testing.T) {
		game := newGameMock()
		game.player1.herdSize = 1
		game.player1.cards = GoodMap{GoodCloth: 1}
		game.cardsOnTable = GoodMap{GoodDiamond: 3}

		error := game.Exchange(GoodMap{GoodDiamond: 3}, GoodMap{GoodCloth: 1, GoodCamel: 2})

		assert.EqualError(t, error, NotEnoughCardsToSellError.Error())
	})

	t.Run("Error if there are not enough cards on the table", func(t *testing.T) {
		game := newGameMock()
		game.cardsOnTable = GoodMap{}

		error := game.Exchange(
			GoodMap{GoodDiamond: 2, GoodGold: 1},
			GoodMap{GoodCloth: 2, GoodLeather: 1},
		)

		assert.EqualError(t, error, NotEnoughCardsOnTableError.Error())
	})

	t.Run("Error if player wants to exchange different amounts", func(t *testing.T) {
		game := newGameMock()
		game.cardsOnTable = GoodMap{GoodCloth: 3}
		game.currentPlayer.cards = GoodMap{GoodDiamond: 2}

		error := game.Exchange(GoodMap{GoodCloth: 3}, GoodMap{GoodDiamond: 2})

		assert.EqualError(t, error, GoodsAmountsMismatchError.Error())
	})

	t.Run("Error if player wants to acquire more than 7 goods which aren't camel", func(t *testing.T) {
		game := newGameMock()
		game.cardsOnTable = GoodMap{GoodCloth: 3}
		game.currentPlayer.cards = GoodMap{GoodDiamond: 7}
		game.currentPlayer.herdSize = 3

		error := game.Exchange(GoodMap{GoodCloth: 1}, GoodMap{GoodCamel: 1})

		assert.EqualError(t, PlayerHasTooManyCardsError, error.Error())
	})

	t.Run("Error if round ended", func(t *testing.T) {
		game := newGameMock()
		game.roundEnded = func() bool { return true }

		error := game.Exchange(GoodMap{GoodCloth: 3}, GoodMap{GoodDiamond: 2})

		assert.EqualError(t, error, RoundEndedError.Error())
	})

	t.Run("Error if game ended", func(t *testing.T) {
		game := newGameMock()
		game.gameEnded = func() bool { return true }

		error := game.Exchange(GoodMap{GoodCloth: 3}, GoodMap{GoodDiamond: 2})

		assert.EqualError(t, error, GameEndedError.Error())
	})
}
