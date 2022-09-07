package jaipur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDrawCard(t *testing.T) {
	t.Run("Player has the choosen cards", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), productMap{}}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: productMap{ProductCamel: Amount(2)},
			cardsInPack:  productMap{ProductCloth: Amount(50)},
		}

		game.DrawCard(player1Name, ProductCamel, 2)

		assert.Equal(t, game.players[player1Name].cards, productMap{ProductCamel: Amount(2)})
	})

	t.Run("Cards on table won't contain the choosen cards", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), productMap{}}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: productMap{ProductCamel: Amount(2)},
			cardsInPack:  productMap{ProductCloth: Amount(50)},
		}

		game.DrawCard(player1Name, ProductCamel, 2)

		assert.Equal(t, game.cardsOnTable[ProductCamel], Amount(0))
	})

	t.Run("The same amount of cards will be moved from the pack to the table", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), productMap{}}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: productMap{ProductDiamond: Amount(3)},
			cardsInPack:  productMap{ProductSpice: Amount(5)},
		}

		game.DrawCard(player1Name, ProductDiamond, 3)

		assert.Equal(t, game.cardsOnTable[ProductSpice], Amount(3))
		assert.Equal(t, game.cardsInPack[ProductSpice], Amount(2))
	})

	t.Run("Error if player doesn't exist", func(t *testing.T) {
		game := game{}

		error := game.DrawCard("John", ProductDiamond, 3)

		assert.EqualError(t, error, PlayerNotExistsError.Error())
	})

	// TODO: implement
	t.Run("Error if not enough cards on table", func(t *testing.T) {
	})
}
