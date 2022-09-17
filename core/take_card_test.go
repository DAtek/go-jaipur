package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeCard(t *testing.T) {
	t.Run("Player has the choosen card", func(t *testing.T) {
		game := newGame()
		game.player1.cards = goodMap{GoodDiamond: Amount(1)}

		game.TakeCard(GoodDiamond)

		assert.Equal(t, goodMap{GoodDiamond: Amount(2)}, game.player1.cards)
	})

	t.Run("Player takes all camels", func(t *testing.T) {
		game := newGame()
		game.cardsOnTable = goodMap{GoodCamel: Amount(5)}

		game.TakeCard(GoodCamel)

		assert.Equal(t, Amount(5), game.player1.herdSize)
	})

	t.Run("Cards on table won't contain the picked card", func(t *testing.T) {
		game := newGame()
		game.cardsOnTable = goodMap{GoodGold: Amount(1)}

		game.TakeCard(GoodGold)

		assert.Equal(t, Amount(0), game.cardsOnTable[GoodGold])
	})

	t.Run("The same amount of cards will be moved from the pack to the table", func(t *testing.T) {
		game := newGame()
		game.cardsOnTable = goodMap{GoodDiamond: Amount(3)}
		game.cardsInPack = goodMap{GoodSpice: Amount(5)}

		game.TakeCard(GoodDiamond)

		assert.Equal(t, Amount(1), game.cardsOnTable[GoodSpice])
		assert.Equal(t, Amount(4), game.cardsInPack[GoodSpice])
	})

	cardsOnTableScenarios := []goodMap{
		{},
		{GoodDiamond: 0},
	}

	for _, s := range cardsOnTableScenarios {
		t.Run("Error if not enough cards on table", func(t *testing.T) {
			game := newGame()
			game.cardsOnTable = s

			e := game.TakeCard(GoodDiamond)

			assert.EqualError(t, e, NotEnoughCardsOnTableError.Error())
		})
	}
}
