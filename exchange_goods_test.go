package jaipur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchangeGoods(t *testing.T) {
	t.Run("Player has the wanted cards", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			cardsInPack:  goodMap{GoodCloth: Amount(50)},
		}

		error := game.ExchangeGoods(
			player1Name,
			goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.Nil(t, error)
		assert.Equal(t, player1.cards[GoodDiamond], Amount(2))
		assert.Equal(t, player1.cards[GoodGold], Amount(1))
	})

	t.Run("Player doesn't have the sold cards", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			cardsInPack:  goodMap{GoodCloth: Amount(50)},
		}

		game.ExchangeGoods(
			player1Name,
			goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.Equal(t, player1.cards[GoodCloth], Amount(0))
		assert.Equal(t, player1.cards[GoodLeather], Amount(0))
	})

	t.Run("Table doesn't have the picked cards", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			cardsInPack:  goodMap{GoodCloth: Amount(50)},
		}

		game.ExchangeGoods(
			player1Name,
			goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.Equal(t, game.cardsOnTable[GoodDiamond], Amount(0))
		assert.Equal(t, game.cardsOnTable[GoodGold], Amount(0))
	})

	t.Run("Table has the dropped cards", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			cardsInPack:  goodMap{GoodCloth: Amount(50)},
		}

		game.ExchangeGoods(
			player1Name,
			goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.Equal(t, game.cardsOnTable[GoodCloth], Amount(2))
		assert.Equal(t, game.cardsOnTable[GoodLeather], Amount(1))
	})

	t.Run("Error if player doesn't exist", func(t *testing.T) {
		game := game{}

		error := game.ExchangeGoods("John", goodMap{}, goodMap{})

		assert.EqualError(t, error, PlayerNotExistsError.Error())
	})

	t.Run("Error if player doesn't have proper cards", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodCloth: 2},
		}

		error := game.ExchangeGoods(player1Name, goodMap{GoodCloth: 2}, goodMap{GoodDiamond: 2})

		assert.EqualError(t, error, NotEnoughCardsToSellError.Error())
	})

	t.Run("Error if there are not enough cards on the table", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players: players,
		}

		error := game.ExchangeGoods(
			player1Name,
			goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
			goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)},
		)

		assert.EqualError(t, error, NotEnoughCardsOnTableError.Error())
	})

	t.Run("Error if player wants to exchange different amounts", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{GoodDiamond: 2}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodCloth: 3},
		}

		error := game.ExchangeGoods(player1Name, goodMap{GoodCloth: 3}, goodMap{GoodDiamond: 2})

		assert.EqualError(t, error, GoodsAmountsMismatchError.Error())
	})
}
