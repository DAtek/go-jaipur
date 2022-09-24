package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSell(t *testing.T) {
	simpleScenarios := []struct {
		name      string
		goodsType GoodType
		amount    Amount
		score     Score
	}{
		{"Sell 2 silvers", GoodSilver, 2, Score(10)},
		{"Sell 2 golds", GoodGold, 2, Score(12)},
		{"Sell 2 cloths", GoodCloth, 2, Score(8)},
		{"Sell 1 spice", GoodSpice, 1, Score(5)},
	}

	for _, s := range simpleScenarios {
		t.Run(s.name, func(t *testing.T) {
			game := newGame()
			game.player1.cards = GoodMap{s.goodsType: s.amount}

			game.Sell(s.goodsType)

			assert.Equal(t, s.score, game.player1.score)
		})
	}

	bonusScenarios := []struct {
		name      string
		goodsType GoodType
		amount    Amount
		minScore  Score
		maxScore  Score
	}{
		{"Sell 3 silvers", GoodSilver, 3, Score(16), Score(18)},
		{"Sell 4 diamonds", GoodDiamond, 4, Score(28), Score(32)},
		{"Sell 5 spices", GoodSpice, 5, Score(20), Score(22)},
		{"Sell 6 spices", GoodSpice, 6, Score(21), Score(23)},
	}

	for _, s := range bonusScenarios {
		t.Run(s.name, func(t *testing.T) {
			game := newGame()
			game.player1.cards = GoodMap{s.goodsType: s.amount}

			game.Sell(s.goodsType)

			assert.GreaterOrEqual(t, game.player1.score, s.minScore)
			assert.LessOrEqual(t, game.player1.score, s.maxScore)
		})
	}

	notEnoughCardsToSellScenarios := []struct {
		name      string
		goodsType GoodType
		amount    Amount
	}{
		{"Sell 1 diamond", GoodDiamond, 1},
		{"Sell 1 gold", GoodGold, 1},
		{"Sell 1 silver", GoodSilver, 1},
		{"Sell 0 leather", GoodLeather, 0},
	}

	for _, s := range notEnoughCardsToSellScenarios {
		t.Run(s.name, func(t *testing.T) {
			game := newGame()
			game.player1.cards = GoodMap{s.goodsType: s.amount}

			error := game.Sell(s.goodsType)

			assert.EqualError(t, error, NotEnoughCardsToSellError.Error())
		})
	}

	t.Run("New score adds up to player's score", func(t *testing.T) {
		game := newGame()
		game.player1.cards = GoodMap{GoodSilver: 2}
		game.player1.score = Score(1)

		game.Sell(GoodSilver)

		assert.Equal(t, Score(11), game.player1.score)
	})

	t.Run("Can't sell camel", func(t *testing.T) {
		game := newGame()
		game.player1.cards = GoodMap{GoodCamel: 2}

		error := game.Sell(GoodCamel)

		assert.EqualError(t, error, SellingCamelForbiddenError.Error())
	})

	t.Run("Current player changes", func(t *testing.T) {
		game := newGame()
		game.player1.cards = GoodMap{GoodDiamond: 2}

		game.Sell(GoodDiamond)

		assert.Equal(t, game.player2.name, game.currentPlayer.name)
	})

	t.Run("Error if round ended", func(t *testing.T) {
		game := newGame()
		*game.roundEnded = func() bool { return true }

		error := game.Sell(GoodCloth)

		assert.EqualError(t, error, RoundEndedError.Error())
	})

	t.Run("Error if game ended", func(t *testing.T) {
		game := newGame()
		*game.gameEnded = func() bool { return true }

		error := game.Sell(GoodCloth)

		assert.EqualError(t, error, GameEndedError.Error())
	})
}
