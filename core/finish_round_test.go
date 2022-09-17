package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinishRound(t *testing.T) {
	excellenceScenarios := []struct {
		score1      Score
		score2      Score
		herd1       Amount
		herd2       Amount
		excellence1 Score
		excellence2 Score
	}{
		{score1: 1, score2: 0, excellence1: 1, excellence2: 1, herd1: 0, herd2: 0},
		{score1: 0, score2: 1, excellence1: 0, excellence2: 2, herd1: 0, herd2: 0},
		{score1: 2, score2: 2, excellence1: 0, excellence2: 1, herd1: 0, herd2: 0},
		{score1: 9, score2: 5, excellence1: 0, excellence2: 2, herd1: 0, herd2: 1},
		{score1: 5, score2: 9, excellence1: 1, excellence2: 1, herd1: 1, herd2: 0},
	}
	for _, s := range excellenceScenarios {
		t.Run("Round winner's excellence points will increase", func(t *testing.T) {
			game := newGame()
			game.player1.score = s.score1
			game.player2.score = s.score2
			game.player1.sealsOfExcellence = 0
			game.player2.sealsOfExcellence = 1
			game.player1.herdSize = s.herd1
			game.player2.herdSize = s.herd2
			game.roundEnded = func() bool { return true }

			game.FinishRound()

			assert.Equal(t, s.excellence1, game.player1.sealsOfExcellence)
			assert.Equal(t, s.excellence2, game.player2.sealsOfExcellence)
		})
	}

	t.Run("Game will be reset", func(t *testing.T) {
		resetCalled := false
		game := newGame()
		game.resetAfterRound = func() {
			resetCalled = true
		}
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
