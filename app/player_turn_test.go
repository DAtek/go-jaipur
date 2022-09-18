package app

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoPlayerAction(t *testing.T) {
	t.Run("Asks for action", func(t *testing.T) {
		m := newMockApp()

		doPlayerAction(m.app)

		assert.True(t, strings.Contains(m.writer.String(), "Pick an action - (E)xchnge cards | (S)ell cards | (T)ake a card: "))
	})

	t.Run("Displays current player's name", func(t *testing.T) {
		m := newMockApp()

		doPlayerAction(m.app)
		wantedString := string(m.app.game.CurrentPlayerName()) + ", it's your turn"

		assert.True(t, strings.Contains(m.writer.String(), wantedString))
	})
}
