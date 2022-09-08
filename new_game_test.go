package jaipur

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	t.Run("Players have correct name", func(t *testing.T) {
		player1 := Name("John")
		player2 := Name("Jane")

		game, _ := NewGame(player1, player2)
		players := game.players

		assert.Equal(t, players[player1].name, player1)
		assert.Equal(t, players[player2].name, player2)
	})

	t.Run("Players have 0 points", func(t *testing.T) {
		player1 := Name("John")
		player2 := Name("Jane")

		game, _ := NewGame(player1, player2)
		players := game.players

		assert.Equal(t, players[player1].score, Score(0))
		assert.Equal(t, players[player2].score, Score(0))
	})

	t.Run("Players have 0 points", func(t *testing.T) {
		player1 := Name("John")
		player2 := Name("Jane")

		game, _ := NewGame(player1, player2)
		players := game.players

		assert.Equal(t, players[player1].cards, productMap{})
		assert.Equal(t, players[player2].cards, productMap{})
	})

	t.Run("Players have 0 camel", func(t *testing.T) {
		player1 := Name("John")
		player2 := Name("Jane")

		game, _ := NewGame(player1, player2)
		players := game.players

		assert.Equal(t, players[player1].herdSize, Amount(0))
		assert.Equal(t, players[player2].herdSize, Amount(0))
	})

	t.Run("There are 5 cards on the table", func(t *testing.T) {
		game, _ := NewGame("a", "b")

		cardsOnTable := Amount(0)
		for _, value := range game.cardsOnTable {
			cardsOnTable += value
		}
		assert.Equal(t, cardsOnTable, Amount(5))
	})

	t.Run("There are at least 3 camels on the table", func(t *testing.T) {
		game, _ := NewGame("a", "b")

		assert.GreaterOrEqual(t, game.cardsOnTable[ProductCamel], Amount(3))
	})

	t.Run("Can't create game with same player names", func(t *testing.T) {
		_, err := NewGame("John", "John")

		assert.EqualError(t, err, SameNamesError.Error())
	})
}
