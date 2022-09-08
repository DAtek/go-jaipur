package jaipur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawCard(t *testing.T) {
	t.Run("Player has the choosen card", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), productMap{ProductDiamond: 1}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: productMap{ProductDiamond: Amount(2)},
			cardsInPack:  productMap{ProductCloth: Amount(50)},
		}

		game.DrawCard(player1Name, ProductDiamond)

		assert.Equal(t, game.players[player1Name].cards, productMap{ProductDiamond: Amount(2)})
	})

	t.Run("Player grabs all camels", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), productMap{}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: productMap{ProductCamel: Amount(5)},
			cardsInPack:  productMap{ProductCloth: Amount(50)},
		}

		game.DrawCard(player1Name, ProductCamel)

		assert.Equal(t, game.players[player1Name].herdSize, Amount(5))
	})

	t.Run("Cards on table won't contain the picked card", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), productMap{}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: productMap{ProductGold: Amount(1)},
			cardsInPack:  productMap{ProductCloth: Amount(50)},
		}

		game.DrawCard(player1Name, ProductGold)

		assert.Equal(t, game.cardsOnTable[ProductGold], Amount(0))
	})

	t.Run("The same amount of cards will be moved from the pack to the table", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), productMap{}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: productMap{ProductDiamond: Amount(3)},
			cardsInPack:  productMap{ProductSpice: Amount(5)},
		}

		game.DrawCard(player1Name, ProductDiamond)

		assert.Equal(t, game.cardsOnTable[ProductSpice], Amount(1))
		assert.Equal(t, game.cardsInPack[ProductSpice], Amount(4))
	})

	t.Run("Error if player doesn't exist", func(t *testing.T) {
		game := game{}

		error := game.DrawCard("John", ProductDiamond)

		assert.EqualError(t, error, PlayerNotExistsError.Error())
	})

	cardsOnTableScenarios := []productMap{
		{},
		{ProductDiamond: 0},
	}

	for _, s := range cardsOnTableScenarios {
		t.Run("Error if not enough cards on table", func(t *testing.T) {
			player1Name := Name("Max")
			player1 := player{player1Name, Score(0), productMap{}, 0}
			players := playerMap{player1Name: &player1}
			game := game{
				players:      players,
				cardsOnTable: s,
			}

			e := game.DrawCard(player1Name, ProductDiamond)

			assert.EqualError(t, e, NotEnoughCardsOnTableError.Error())
		})
	}
}
