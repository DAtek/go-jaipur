package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchangeGoods(t *testing.T) {
	t.Run("Player has the wanted cards", func(t *testing.T) {
		game := newGame()
		game.currentPlayer.cards = GoodMap{
			GoodCloth:   Amount(2),
			GoodLeather: Amount(1),
		}

		error := game.ExchangeGoods(
			GoodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.Nil(t, error)
		assert.Equal(t, Amount(2), game.player1.cards[GoodDiamond])
		assert.Equal(t, Amount(1), game.player1.cards[GoodGold])
	})

	t.Run("Player doesn't have the sold cards", func(t *testing.T) {
		game := newGame()

		game.ExchangeGoods(
			GoodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.Equal(t, Amount(0), game.player1.cards[GoodCloth])
		assert.Equal(t, Amount(0), game.player1.cards[GoodLeather])
	})

	t.Run("Table doesn't have the picked cards", func(t *testing.T) {
		game := newGame()

		game.ExchangeGoods(
			GoodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.Equal(t, Amount(0), game.cardsOnTable[GoodDiamond])
		assert.Equal(t, Amount(0), game.cardsOnTable[GoodGold])
	})

	t.Run("Table has the dropped cards", func(t *testing.T) {
		game := newGame()

		game.ExchangeGoods(
			GoodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.Equal(t, Amount(2), game.cardsOnTable[GoodCloth])
		assert.Equal(t, Amount(1), game.cardsOnTable[GoodLeather])
	})

	t.Run("Current player changes", func(t *testing.T) {
		game := newGame()

		game.ExchangeGoods(
			GoodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.Equal(t, game.player2.name, game.currentPlayer.name)
	})

	t.Run("Error if player doesn't have proper cards", func(t *testing.T) {
		game := newGame()
		game.player1.cards = GoodMap{}
		game.cardsOnTable = GoodMap{GoodCloth: 2}

		error := game.ExchangeGoods(GoodMap{GoodCloth: 2}, GoodMap{GoodDiamond: 2})

		assert.EqualError(t, error, NotEnoughCardsToSellError.Error())
	})

	t.Run("Error if there are not enough cards on the table", func(t *testing.T) {
		game := newGame()
		game.cardsOnTable = GoodMap{}

		error := game.ExchangeGoods(
			GoodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.EqualError(t, error, NotEnoughCardsOnTableError.Error())
	})

	t.Run("Error if player wants to exchange different amounts", func(t *testing.T) {
		game := newGame()
		game.cardsOnTable = GoodMap{GoodCloth: 3}
		game.currentPlayer.cards = GoodMap{GoodDiamond: 2}

		error := game.ExchangeGoods(GoodMap{GoodCloth: 3}, GoodMap{GoodDiamond: 2})

		assert.EqualError(t, error, GoodsAmountsMismatchError.Error())
	})

	t.Run("Error if round ended", func(t *testing.T) {
		game := newGame()
		game.roundEnded = func() bool { return true }

		error := game.ExchangeGoods(GoodMap{GoodCloth: 3}, GoodMap{GoodDiamond: 2})

		assert.EqualError(t, error, RoundEndedError.Error())
	})

	t.Run("Error if game ended", func(t *testing.T) {
		game := newGame()
		game.gameEnded = func() bool { return true }

		error := game.ExchangeGoods(GoodMap{GoodCloth: 3}, GoodMap{GoodDiamond: 2})

		assert.EqualError(t, error, GameEndedError.Error())
	})
}
