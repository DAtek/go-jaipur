package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSellGoods(t *testing.T) {
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
			player1Name := Name("John")
			player1 := player{player1Name, Score(0), goodMap{s.goodsType: s.amount}, 0}
			players := playerMap{player1Name: &player1}

			game := game{
				players:   players,
				soldGoods: goodMap{},
			}

			game.SellGoods(&player1Name, s.goodsType)

			assert.Equal(t, player1.score, s.score)
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
			player1Name := Name("John")
			player1 := player{player1Name, Score(0), goodMap{s.goodsType: s.amount}, 0}
			players := playerMap{player1Name: &player1}

			game := game{
				players:   players,
				soldGoods: goodMap{},
			}

			game.SellGoods(&player1Name, s.goodsType)

			assert.GreaterOrEqual(t, player1.score, s.minScore)
			assert.LessOrEqual(t, player1.score, s.maxScore)
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
			player1Name := Name("John")
			player1 := player{player1Name, Score(0), goodMap{s.goodsType: s.amount}, 0}
			players := playerMap{player1Name: &player1}

			game := game{
				players:   players,
				soldGoods: goodMap{},
			}

			error := game.SellGoods(&player1Name, s.goodsType)

			assert.EqualError(t, error, NotEnoughCardsToSellError.Error())
		})
	}

	t.Run("Error if player doesn't exists", func(t *testing.T) {
		game := game{
			players:   playerMap{},
			soldGoods: goodMap{},
		}

		fakePlayerName := Name("a")

		error := game.SellGoods(&fakePlayerName, GoodSilver)
		assert.EqualError(t, error, PlayerNotExistsError.Error())
	})

	t.Run("Can't sell camel", func(t *testing.T) {
		player1Name := Name("John")
		player1 := player{player1Name, Score(0), goodMap{GoodCamel: 2}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:   players,
			soldGoods: goodMap{},
		}

		error := game.SellGoods(&player1Name, GoodCamel)

		assert.EqualError(t, error, SellingCamelForbiddenError.Error())
	})
}
