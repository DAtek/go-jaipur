package core

import (
	"testing"

	"github.com/DAtek/gotils"
	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	playerNames := [][]Name{
		{"", "a"},
		{"a", ""},
	}
	for _, s := range playerNames {
		t.Run("Players can't have empty name", func(t *testing.T) {
			player1 := Name(s[0])
			player2 := Name(s[1])

			_, err := newGame(player1, player2)

			assert.EqualError(t, EmptyNameError, err.Error())
		})
	}

	t.Run("Players have correct name", func(t *testing.T) {
		player1 := Name("a")
		player2 := Name("b")

		game := gotils.ResultOrPanic(newGame(player1, player2))

		assert.Equal(t, player1, game.player1.name)
		assert.Equal(t, player2, game.player2.name)
	})

	t.Run("Current player is player 1", func(t *testing.T) {
		player1 := Name("a")
		player2 := Name("b")

		game, _ := newGame(player1, player2)

		assert.Equal(t, player1, game.currentPlayer.name)
	})

	t.Run("Players have 0 points", func(t *testing.T) {
		player1 := Name("a")
		player2 := Name("b")

		game, _ := newGame(player1, player2)

		assert.Equal(t, Score(0), game.player1.score)
		assert.Equal(t, Score(0), game.player2.score)
	})

	t.Run("Players have 0 seals of excellence", func(t *testing.T) {
		player1 := Name("a")
		player2 := Name("b")

		game, _ := newGame(player1, player2)

		assert.Equal(t, Score(0), game.player1.sealsOfExcellence)
		assert.Equal(t, Score(0), game.player2.sealsOfExcellence)
	})

	t.Run("There are 5 cards on the table", func(t *testing.T) {
		game, _ := newGame("a", "b")

		cardsOnTable := Amount(0)
		for _, value := range game.cardsOnTable {
			cardsOnTable += value
		}
		assert.Equal(t, Amount(5), cardsOnTable)
	})

	t.Run("There are at least 3 camels on the table", func(t *testing.T) {
		game, _ := newGame("a", "b")

		assert.GreaterOrEqual(t, game.cardsOnTable[GoodCamel], Amount(3))
	})

	t.Run("There are 40 cards in the pack", func(t *testing.T) {
		game, _ := newGame("a", "b")

		cardsInPack := Amount(0)
		for _, amount := range game.cardsInPack {
			cardsInPack += amount
		}

		assert.Equal(t, Amount(40), cardsInPack)
	})

	t.Run("Can't create game with same player names", func(t *testing.T) {
		_, err := NewGame("a", "a")

		assert.Error(t, err)
	})

	t.Run("Round ended method is working", func(t *testing.T) {
		game, _ := NewGame("a", "b")

		assert.False(t, game.RoundEnded())
	})

	t.Run("Game ended method is working", func(t *testing.T) {
		game, _ := NewGame("a", "b")

		assert.False(t, game.GameEnded())
	})

	t.Run("Round winner method is working", func(t *testing.T) {
		game, _ := NewGame("a", "b")

		_, err := game.RoundWinner()

		assert.Error(t, err)
	})
}

func newGame(player1, player2 Name) (*game, error) {
	igame, err := NewGame(player1, player2)
	if igame == nil {
		return nil, err
	}

	game, ok := igame.(*game)
	if !ok {
		panic("COULD_NOT_CONVERT_IGAME_TO_GAME")
	}
	return game, err
}
