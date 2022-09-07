package jaipur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSellProduct(t *testing.T) {
	simpleScenarios := []struct {
		name      string
		goodsType ProductType
		amount    Amount
		score     Score
	}{
		{"Sell 2 silvers", ProductSilver, 2, Score(10)},
		{"Sell 2 golds", ProductGold, 2, Score(12)},
		{"Sell 2 cloths", ProductCloth, 2, Score(8)},
	}

	for _, s := range simpleScenarios {
		t.Run(s.name, func(t *testing.T) {
			player1Name := Name("John")
			player1 := player{player1Name, Score(0), productMap{s.goodsType: s.amount}}
			players := playerMap{player1Name: &player1}

			game := game{
				players:      players,
				soldProducts: productMap{},
			}

			game.SellProduct(&player1Name, s.goodsType, s.amount)

			assert.Equal(t, player1.score, s.score)
		})
	}

	bonusScenarios := []struct {
		name      string
		goodsType ProductType
		amount    Amount
		minScore  Score
		maxScore  Score
	}{
		{"Sell 3 silvers", ProductSilver, 3, Score(16), Score(18)},
		{"Sell 4 diamonds", ProductDiamond, 4, Score(28), Score(32)},
		{"Sell 5 spices", ProductSpice, 5, Score(20), Score(22)},
		{"Sell 6 spices", ProductSpice, 6, Score(21), Score(23)},
	}

	for _, s := range bonusScenarios {
		t.Run(s.name, func(t *testing.T) {
			player1Name := Name("John")
			player1 := player{player1Name, Score(0), productMap{s.goodsType: s.amount}}
			players := playerMap{player1Name: &player1}

			game := game{
				players:      players,
				soldProducts: productMap{},
			}

			game.SellProduct(&player1Name, s.goodsType, s.amount)

			assert.GreaterOrEqual(t, player1.score, s.minScore)
			assert.LessOrEqual(t, player1.score, s.maxScore)
		})
	}

	t.Run("Error if not enough cards to sell", func(t *testing.T) {
		player1Name := Name("John")
		player1 := player{player1Name, Score(0), productMap{}}
		players := playerMap{player1Name: &player1}

		game := game{
			players:      players,
			soldProducts: productMap{},
		}

		error := game.SellProduct(&player1Name, ProductSilver, 1)

		assert.EqualError(t, error, NotEnoughCardsError.Error())
	})

	t.Run("Error if player doesn't exists", func(t *testing.T) {
		game := game{
			players:      playerMap{},
			soldProducts: productMap{},
		}

		fakePlayerName := Name("a")

		error := game.SellProduct(&fakePlayerName, ProductSilver, 1)
		assert.EqualError(t, error, PlayerNotExistsError.Error())
	})
}
