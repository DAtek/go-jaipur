package app

import (
	"jaipur/core"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTakeCard(t *testing.T) {
	t.Run("Asks card for taking a card", func(t *testing.T) {
		mockApp := newMockApp()

		takeCard(mockApp.app)

		assert.True(t, strings.Contains(mockApp.writer.String(), "Take a card: "))
	})

	t.Run("Transition clears screen on success", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.reader.Write([]byte("D"))

		takeCard(mockApp.app)

		assert.True(t, strings.Contains(mockApp.writer.String(), clearScreenString))
	})

	t.Run("Turn ended when there are no more cards", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.game.roundEnded = true
		mockApp.reader.Write([]byte("D"))

		nextState := takeCard(mockApp.app)

		assert.Equal(t, roundEnded.Name, nextState)
	})

	t.Run("Transition called take card", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.game.roundEnded = true
		called := false
		mockApp.game.takeCard = func(card core.GoodType) error {
			called = true
			return nil
		}
		mockApp.reader.Write([]byte("D"))

		takeCard(mockApp.app)

		assert.True(t, called)
	})

	t.Run("Transition prints error on wrong input", func(t *testing.T) {
		mockApp := newMockApp()
		called := false
		mockApp.game.takeCard = func(card core.GoodType) error {
			called = true
			return nil
		}
		mockApp.reader.Write([]byte("asdasdasd"))

		nextState := takeCard(mockApp.app)

		assert.False(t, called)
		assert.True(t, strings.Contains(mockApp.writer.String(), "Invalid input"))
		assert.Equal(t, playerTurn.Name, nextState)
	})

	t.Run("Transition prints game logic error", func(t *testing.T) {
		mockApp := newMockApp()
		called := false
		mockApp.game.takeCard = func(card core.GoodType) error {
			called = true
			return core.PlayerHasTooManyCardsError
		}
		mockApp.reader.Write([]byte("D"))

		nextState := takeCard(mockApp.app)

		assert.True(t, called)
		assert.True(t, strings.Contains(mockApp.writer.String(), core.PlayerHasTooManyCardsError.Error()))
		assert.Equal(t, playerTurn.Name, nextState)
	})
}
