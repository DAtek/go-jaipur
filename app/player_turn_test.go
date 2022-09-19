package app

import (
	"jaipur/fsm"
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDoPlayerAction(t *testing.T) {
	t.Run("Asks for action", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := newMockPlayerCommand()

		doPlayerAction(mockApp.app, playerCommands)

		assert.True(t, strings.Contains(mockApp.writer.String(), "Pick an action - (E)xchnge cards | (S)ell cards | (T)ake a card: "))
	})

	t.Run("Displays current player's name", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := newMockPlayerCommand()

		doPlayerAction(mockApp.app, playerCommands)
		wantedString := string(mockApp.app.game.CurrentPlayerName()) + ", it's your turn"

		assert.True(t, strings.Contains(mockApp.writer.String(), wantedString))
	})

	t.Run("Displays current player's cards", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := newMockPlayerCommand()

		doPlayerAction(mockApp.app, playerCommands)
		wantedString := "Your cards: " + formatGoodMap(mockApp.app.game.CurrentPlayerCards())

		assert.True(t, strings.Contains(mockApp.writer.String(), wantedString))
	})

	t.Run("Displays cards on table", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := newMockPlayerCommand()

		doPlayerAction(mockApp.app, playerCommands)
		wantedString := "Cards on table: " + formatGoodMap(mockApp.app.game.CardsOnTable())

		assert.True(t, strings.Contains(mockApp.writer.String(), wantedString))
	})

	t.Run("Prints warning on wrong command", func(t *testing.T) {
		mockApp := newMockApp()
		playerCommands := newMockPlayerCommand()
		mockApp.reader.Write([]byte("wrong command"))

		doPlayerAction(mockApp.app, playerCommands)

		assert.True(t, strings.Contains(mockApp.writer.String(), "Wrong action"))
	})

	playerChoiceScenarios := []struct {
		key     string
		command string
	}{
		{"T", "TakeCard"},
		{"t", "TakeCard"},
		{"E", "ExchangeCards"},
		{"e", "ExchangeCards"},
		{"S", "SellCards"},
		{"s", "SellCards"},
	}

	for _, s := range playerChoiceScenarios {
		t.Run("Executes the proper command", func(t *testing.T) {
			mockApp := newMockApp()
			playerCommands := newMockPlayerCommand()

			called := false
			nextState := fsm.StateName("next state")
			command := func() fsm.StateName {
				called = true
				return nextState
			}
			structValue := reflect.ValueOf(playerCommands).Elem()
			structFieldValue := structValue.FieldByName(s.command)
			structFieldValue.Set(reflect.ValueOf(command))

			mockApp.reader.Write([]byte(s.key))

			result := doPlayerAction(mockApp.app, playerCommands)

			assert.Equal(t, nextState, result)
			assert.True(t, called)
		})
	}
}
