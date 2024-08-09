package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuy(t *testing.T) {
	t.Run("Player has the choosen card", func(t *testing.T) {
		game := newGameMock()
		game.player1.cards = GoodMap{GoodDiamond: Amount(1)}

		game.Buy(GoodDiamond)

		assert.Equal(t, GoodMap{GoodDiamond: Amount(2)}, game.player1.cards)
	})

	t.Run("Current player changes", func(t *testing.T) {
		game := newGameMock()
		game.player2.cards = GoodMap{GoodDiamond: Amount(1)}
		game.currentPlayer = game.player2

		game.Buy(GoodDiamond)

		assert.Equal(t, game.player1, game.currentPlayer)
	})

	t.Run("Player takes all camels", func(t *testing.T) {
		game := newGameMock()
		game.cardsOnTable = GoodMap{GoodCamel: Amount(5)}

		game.Buy(GoodCamel)

		assert.Equal(t, Amount(5), game.player1.herdSize)
	})

	t.Run("Cards on table won't contain the picked card", func(t *testing.T) {
		game := newGameMock()
		game.cardsOnTable = GoodMap{GoodGold: Amount(1)}

		game.Buy(GoodGold)

		assert.Equal(t, Amount(0), game.cardsOnTable[GoodGold])
	})

	t.Run("The same amount of cards will be moved from the pack to the table", func(t *testing.T) {
		game := newGameMock()
		game.cardsOnTable = GoodMap{GoodDiamond: Amount(3)}
		game.cardsInPack = GoodMap{GoodSpice: Amount(5)}

		game.Buy(GoodDiamond)

		assert.Equal(t, Amount(1), game.cardsOnTable[GoodSpice])
		assert.Equal(t, Amount(4), game.cardsInPack[GoodSpice])
	})

	cardsOnTableScenarios := []GoodMap{
		{},
		{GoodDiamond: 0},
	}

	for _, s := range cardsOnTableScenarios {
		t.Run("Error if not enough cards on table", func(t *testing.T) {
			game := newGameMock()
			game.cardsOnTable = s

			e := game.Buy(GoodDiamond)

			assert.EqualError(t, e, NotEnoughCardsOnTableError.Error())
		})
	}

	t.Run("There is no enough cards in pack", func(t *testing.T) {
		game := newGameMock()
		game.cardsOnTable = GoodMap{GoodDiamond: Amount(5)}
		game.cardsInPack = GoodMap{}

		game.Buy(GoodDiamond)

		cards := Amount(0)
		for _, amount := range game.cardsOnTable {
			cards += amount
		}

		assert.Equal(t, Amount(4), cards)
	})

	t.Run("Error if round ended", func(t *testing.T) {
		game := newGameMock()
		game.roundEnded = func() bool { return true }

		error := game.Buy(GoodCloth)

		assert.EqualError(t, error, RoundEndedError.Error())
	})

	t.Run("Error if game ended", func(t *testing.T) {
		game := newGameMock()
		game.gameEnded = func() bool { return true }

		error := game.Buy(GoodCloth)

		assert.EqualError(t, error, GameEndedError.Error())
	})

	t.Run("Error if player has too many cards", func(t *testing.T) {
		game := newGameMock()
		game.player1.cards = GoodMap{GoodDiamond: 7}

		error := game.Buy(GoodCloth)

		assert.EqualError(t, error, PlayerHasTooManyCardsError.Error())
	})
}
