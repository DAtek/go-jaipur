package app

import (
	"fmt"
	"jaipur/core"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFinishRound(t *testing.T) {
	t.Run("Prints the winner of the round and the scores", func(t *testing.T) {
		mockApp := newMockApp()
		winner := core.Name("Susie")
		roundWinner := func() (core.Name, error) {
			return winner, nil
		}
		mockApp.game.roundWinner = &roundWinner
		susie := core.Name("Susie")
		johann := core.Name("Johann")
		mockApp.game.playerScores = core.ScoreMap{
			core.Name(susie):  4,
			core.Name(johann): 2,
		}

		finishRound(mockApp.app)

		output := mockApp.writer.String()
		fmt.Printf("output: %s\n", output)
		assert.True(t, strings.Contains(output, "Winner of the round: "+string(winner)))
		assert.True(t, strings.Contains(output, fmt.Sprintf("%s's score: %d", susie, mockApp.game.playerScores[susie])))
		assert.True(t, strings.Contains(output, fmt.Sprintf("%s's score: %d", johann, mockApp.game.playerScores[johann])))
	})

	t.Run("Prompts to continue", func(t *testing.T) {
		mockApp := newMockApp()

		finishRound(mockApp.app)

		output := mockApp.writer.String()

		assert.True(t, strings.Contains(output, "Press enter to continue"))
		assert.True(t, strings.Contains(output, clearScreenString))
	})

	t.Run("Calls core finish round", func(t *testing.T) {
		mockApp := newMockApp()
		called := false
		mockApp.game.finishRound = func() error {
			called = true
			return nil
		}

		nextState := finishRound(mockApp.app)

		assert.True(t, called)
		assert.Equal(t, playerTurn.Name, nextState)
	})

	t.Run("Panics when finish round return error", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.game.finishRound = func() error {
			return core.RoundNotEndedError
		}

		assert.Panics(t, func() { finishRound(mockApp.app) })
	})

	t.Run("Next state is game ended when the game is ended", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.game.gameEnded = true

		nextState := finishRound(mockApp.app)

		assert.Equal(t, gameEnded.Name, nextState)
	})
}
