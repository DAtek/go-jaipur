package app

import (
	"jaipur/core"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAskForNames(t *testing.T) {
	t.Run("Returns original state when players have the same name", func(t *testing.T) {
		m := newMockApp()
		m.reader.Write([]byte("a"))
		m.reader.Write([]byte("\n"))
		m.reader.Write([]byte("a"))

		nextState := askForNames(m.app)
		output := m.writer.String()

		assert.Equal(t, gameStart.Name, nextState)
		assert.True(t, strings.Contains(output, core.SameNamesError.Error()))
	})

	t.Run("Collects player names", func(t *testing.T) {
		player1 := "a"
		player2 := "b"
		m := newMockApp()
		m.reader.Write([]byte(player1))
		m.reader.Write([]byte("\n"))
		m.reader.Write([]byte(player2))

		nextState := askForNames(m.app)
		output := m.writer.String()

		assert.Equal(t, playerTurn.Name, nextState)
		assert.True(t, strings.Contains(output, clearScreenString))
	})

	t.Run("Asks for player names", func(t *testing.T) {
		m := newMockApp()

		askForNames(m.app)

		output := m.writer.String()

		assert.True(t, strings.Contains(output, "Enter player 1 name: "))
		assert.True(t, strings.Contains(output, "Enter player 2 name: "))
	})
}
