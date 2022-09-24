package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundWinner(t *testing.T) {
	t.Run("Returns round not ended error if round not ended", func(t *testing.T) {
		game := newGame()
		*game.roundEnded = func() bool { return false }

		_, err := roundWinner(game)

		assert.EqualError(t, RoundNotEndedError, err.Error())
	})

	player1Name := Name("John")
	player2Name := Name("Jane")
	winnerScenarios := []struct {
		player1Score Score
		player2Score Score
		player1Herd  Amount
		player2Herd  Amount
		winner       Name
	}{
		{20, 21, 0, 0, player2Name},
		{31, 30, 0, 0, player1Name},
		{5, 9, 2, 1, player1Name},
		{9, 5, 1, 2, player2Name},
	}
	for _, s := range winnerScenarios {
		t.Run("Returns the player with the biggest score", func(t *testing.T) {
			game := newGame()
			*game.roundEnded = func() bool { return true }
			game.player1.score = s.player1Score
			game.player2.score = s.player2Score
			game.player1.name = player1Name
			game.player2.name = player2Name
			game.player1.herdSize = s.player1Herd
			game.player2.herdSize = s.player2Herd

			winner, err := roundWinner(game)

			assert.Nil(t, err)
			assert.Equal(t, s.winner, winner)
		})
	}

	t.Run("No winner if players have the same score", func(t *testing.T) {
		game := newGame()
		*game.roundEnded = func() bool { return true }

		winner, err := roundWinner(game)

		assert.Nil(t, err)
		assert.Equal(t, Name(EmptyName), winner)
	})
}
