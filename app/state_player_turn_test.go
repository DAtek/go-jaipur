package app

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerTurn(t *testing.T) {
	t.Run("Asks for action", func(t *testing.T) {
		mockApp := newMockApp()

		playerTurn.Transit(mockApp.app)

		assert.True(t, strings.Contains(mockApp.writer.String(), "Pick an action - (E)xchange | (S)ell | (B)uy"))
	})

	t.Run("Displays current player's name", func(t *testing.T) {
		mockApp := newMockApp()

		playerTurn.Transit(mockApp.app)
		wantedString := string(mockApp.app.Game.CurrentPlayerName()) + ", it's your turn"

		assert.True(t, strings.Contains(mockApp.writer.String(), wantedString))
	})

	t.Run("Displays current player's cards", func(t *testing.T) {
		mockApp := newMockApp()

		playerTurn.Transit(mockApp.app)
		wantedString := "Your cards: " + formatGoodMap(mockApp.app.Game.CurrentPlayerCards())

		assert.True(t, strings.Contains(mockApp.writer.String(), wantedString))
	})

	t.Run("Displays cards on table", func(t *testing.T) {
		mockApp := newMockApp()

		playerTurn.Transit(mockApp.app)
		wantedString := "Cards on table: " + formatGoodMap(mockApp.app.Game.CardsOnTable())

		assert.True(t, strings.Contains(mockApp.writer.String(), wantedString))
	})

	t.Run("Prints warning on wrong command", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.reader.Write([]byte("wrong command"))

		playerTurn.Transit(mockApp.app)

		assert.True(t, strings.Contains(mockApp.writer.String(), "Wrong action"))
	})
}
