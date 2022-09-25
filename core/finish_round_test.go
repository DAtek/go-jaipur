package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinishRound(t *testing.T) {
	t.Run("Round winner's excellence points will be increased", func(t *testing.T) {
		game := newGame()
		game.player1.sealsOfExcellence = 1
		game.roundEnded = func() bool { return true }
		game.roundWinner = func() (Name, error) {
			return game.player1.name, nil
		}

		game.FinishRound()

		assert.Equal(t, Score(2), game.player1.sealsOfExcellence)
	})

	t.Run("Game will be reset", func(t *testing.T) {
		resetCalled := false
		game := newGame()
		game.resetAfterRound = func() { resetCalled = true }
		game.roundEnded = func() bool { return true }

		game.FinishRound()

		assert.True(t, resetCalled)
	})

	t.Run("Error if round not ended", func(t *testing.T) {
		game := newGame()

		e := game.FinishRound()

		assert.EqualError(t, RoundNotEndedError, e.Error())
	})

	t.Run("Error if game ended", func(t *testing.T) {
		game := newGame()
		game.gameEnded = func() bool { return true }

		error := game.FinishRound()

		assert.EqualError(t, error, GameEndedError.Error())
	})
}
