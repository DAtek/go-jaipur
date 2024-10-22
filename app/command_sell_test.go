package app

import (
	"jaipur/core"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSell(t *testing.T) {
	t.Run("Asks player to sell goods", func(t *testing.T) {
		mockApp := newMockApp()

		sell(mockApp.app)

		assert.True(t, strings.Contains(mockApp.writer.String(), "Pick a good to sell: "))
	})

	t.Run("Calls core sell goods on correct input", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.reader.Write([]byte("G"))
		called := false
		mockApp.game.sell = func(card core.GoodType) error {
			called = true
			return nil
		}

		nextState := sell(mockApp.app)

		assert.True(t, called)
		assert.Equal(t, STATE_PLAYER_TURN, nextState)
	})

	t.Run("Next state is round ended when round ended", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.reader.Write([]byte("G"))
		mockApp.game.roundEnded = true

		nextState := sell(mockApp.app)

		assert.Equal(t, STATE_ROUND_ENDED, nextState)
	})

	t.Run("Next state is player turn when game error happens", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.reader.Write([]byte("G"))
		called := false
		mockApp.game.sell = func(card core.GoodType) error {
			called = true
			return core.SellingCamelForbiddenError
		}

		nextState := sell(mockApp.app)

		assert.Equal(t, STATE_PLAYER_TURN, nextState)
		assert.True(t, called)
	})

	t.Run("Error is logged on game error", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.reader.Write([]byte("G"))
		mockApp.game.sell = func(card core.GoodType) error {
			return core.SellingCamelForbiddenError
		}

		sell(mockApp.app)
		output := mockApp.writer.String()

		assert.True(t, strings.Contains(output, core.SellingCamelForbiddenError.Error()))

	})

	t.Run("Next state is player turn on invalid input", func(t *testing.T) {
		mockApp := newMockApp()
		mockApp.reader.Write([]byte("invalid input"))
		called := false
		mockApp.game.sell = func(card core.GoodType) error {
			called = true
			return nil
		}

		nextState := sell(mockApp.app)
		output := mockApp.writer.String()

		assert.Equal(t, STATE_PLAYER_TURN, nextState)
		assert.True(t, strings.Contains(output, "Invalid input"))
		assert.False(t, called)

	})
}
