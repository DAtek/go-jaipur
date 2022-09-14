package jaipur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeCard(t *testing.T) {
	t.Run("Player has the choosen card", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{GoodDiamond: 1}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodDiamond: Amount(2)},
			cardsInPack:  goodMap{GoodCloth: Amount(50)},
		}

		game.TakeCard(player1Name, GoodDiamond)

		assert.Equal(t, game.players[player1Name].cards, goodMap{GoodDiamond: Amount(2)})
	})

	t.Run("Player takes all camels", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodCamel: Amount(5)},
			cardsInPack:  goodMap{GoodCloth: Amount(50)},
		}

		game.TakeCard(player1Name, GoodCamel)

		assert.Equal(t, game.players[player1Name].herdSize, Amount(5))
	})

	t.Run("Cards on table won't contain the picked card", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodGold: Amount(1)},
			cardsInPack:  goodMap{GoodCloth: Amount(50)},
		}

		game.TakeCard(player1Name, GoodGold)

		assert.Equal(t, game.cardsOnTable[GoodGold], Amount(0))
	})

	t.Run("The same amount of cards will be moved from the pack to the table", func(t *testing.T) {
		player1Name := Name("Max")
		player1 := player{player1Name, Score(0), goodMap{}, 0}
		players := playerMap{player1Name: &player1}
		game := game{
			players:      players,
			cardsOnTable: goodMap{GoodDiamond: Amount(3)},
			cardsInPack:  goodMap{GoodSpice: Amount(5)},
		}

		game.TakeCard(player1Name, GoodDiamond)

		assert.Equal(t, game.cardsOnTable[GoodSpice], Amount(1))
		assert.Equal(t, game.cardsInPack[GoodSpice], Amount(4))
	})

	t.Run("Error if player doesn't exist", func(t *testing.T) {
		game := game{}

		error := game.TakeCard("John", GoodDiamond)

		assert.EqualError(t, error, PlayerNotExistsError.Error())
	})

	cardsOnTableScenarios := []goodMap{
		{},
		{GoodDiamond: 0},
	}

	for _, s := range cardsOnTableScenarios {
		t.Run("Error if not enough cards on table", func(t *testing.T) {
			player1Name := Name("Max")
			player1 := player{player1Name, Score(0), goodMap{}, 0}
			players := playerMap{player1Name: &player1}
			game := game{
				players:      players,
				cardsOnTable: s,
			}

			e := game.TakeCard(player1Name, GoodDiamond)

			assert.EqualError(t, e, NotEnoughCardsOnTableError.Error())
		})
	}
}
