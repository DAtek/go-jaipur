package app

import (
	"jaipur/core"
	"strings"
	"testing"

	"github.com/DAtek/fsm"
	"github.com/stretchr/testify/assert"
)

func TestGameEnded(t *testing.T) {
	t.Run("Displays winner", func(t *testing.T) {
		mockApp := newMockApp()
		name := "Clara"
		gameWinner := func() (core.Name, error) {
			return core.Name(name), nil
		}
		mockApp.game.gameWinner = gameWinner

		gameEnded.Transit(mockApp.app)

		output := mockApp.writer.String()

		assert.True(t, strings.Contains(output, "Congratulations, "+name))
	})

	t.Run("Prompts to continue", func(t *testing.T) {
		mockApp := newMockApp()

		gameEnded.Transit(mockApp.app)

		output := mockApp.writer.String()

		assert.True(t, strings.Contains(output, "Press enter to quit"))
	})

	t.Run("Returns final state", func(t *testing.T) {
		mockApp := newMockApp()

		nextState := gameEnded.Transit(mockApp.app)

		assert.Equal(t, fsm.STATE_FINAL, nextState)
	})

	t.Run("Panics if core game not ended", func(t *testing.T) {
		mockApp := newMockApp()
		gameWinner := func() (core.Name, error) {
			return "", core.GameNotEndedError
		}
		mockApp.game.gameWinner = gameWinner

		assert.Panics(t, func() { gameEnded.Transit(mockApp.app) })
	})
}
