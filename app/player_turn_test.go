package app

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoPlayerAction(t *testing.T) {
	t.Run("Asks for action", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := playerCommandCollection{}

		doPlayerAction(mockApp.app, playerCommands)

		assert.True(t, strings.Contains(mockApp.writer.String(), "Pick an action - (E)xchange | (S)ell | (B)uy"))
	})

	t.Run("Displays current player's name", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := playerCommandCollection{}

		doPlayerAction(mockApp.app, playerCommands)
		wantedString := string(mockApp.app.game.CurrentPlayerName()) + ", it's your turn"

		assert.True(t, strings.Contains(mockApp.writer.String(), wantedString))
	})

	t.Run("Displays current player's cards", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := playerCommandCollection{}

		doPlayerAction(mockApp.app, playerCommands)
		wantedString := "Your cards: " + formatGoodMap(mockApp.app.game.CurrentPlayerCards())

		assert.True(t, strings.Contains(mockApp.writer.String(), wantedString))
	})

	t.Run("Displays cards on table", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := playerCommandCollection{}

		doPlayerAction(mockApp.app, playerCommands)
		wantedString := "Cards on table: " + formatGoodMap(mockApp.app.game.CardsOnTable())

		assert.True(t, strings.Contains(mockApp.writer.String(), wantedString))
	})

	t.Run("Prints warning on wrong command", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := playerCommandCollection{}
		mockApp.reader.Write([]byte("wrong command"))

		doPlayerAction(mockApp.app, playerCommands)

		assert.True(t, strings.Contains(mockApp.writer.String(), "Wrong action"))
	})
}
