package app

import (
	"jaipur/core"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestExchange(t *testing.T) {
	t.Run("Asks for goods to buy", func(t *testing.T) {
		mockApp := newMockApp()

		exchange(mockApp.app, mockApp.exchangeInputParser)
		output := mockApp.writer.String()

		assert.True(t, strings.Contains(output, "Buy goods"))
	})

	t.Run("Asks for goods to sell if the buy input is valid", func(t *testing.T) {
		mockApp := newMockApp()

		exchange(mockApp.app, mockApp.exchangeInputParser)
		output := mockApp.writer.String()

		assert.True(t, strings.Contains(output, "Sell goods"))
	})

	t.Run("Calls core ExchangeGoods", func(t *testing.T) {
		mockApp := newMockApp()
		called := false
		mockApp.game.exchange = func(buy, sell core.GoodMap) error {
			called = true
			return nil
		}

		nextState := exchange(mockApp.app, mockApp.exchangeInputParser)
		output := mockApp.writer.String()

		assert.True(t, called)
		assert.Equal(t, playerTurn.Name, nextState)
		assert.True(t, strings.Contains(output, clearScreenString))
	})

	t.Run("Displays error message on invalid buy input", func(t *testing.T) {
		mockApp := newMockApp()
		called := false
		mockApp.game.exchange = func(buy, sell core.GoodMap) error {
			called = true
			return nil
		}
		mockApp.exchangeInputParser = func(s string) (core.GoodMap, bool) {
			return core.GoodMap{}, false
		}
		mockApp.writer.Write([]byte("oh my"))

		nextState := exchange(mockApp.app, mockApp.exchangeInputParser)
		output := mockApp.writer.String()

		assert.True(t, strings.Contains(output, "Invalid input for buying."))
		assert.False(t, called)
		assert.Equal(t, playerTurn.Name, nextState)
	})

	t.Run("Displays error message on invalid sell input", func(t *testing.T) {
		mockApp := newMockApp()
		called := false
		mockApp.game.exchange = func(buy, sell core.GoodMap) error {
			called = true
			return nil
		}
		inputCounter := 0
		mockApp.exchangeInputParser = func(s string) (core.GoodMap, bool) {
			inputCounter++
			success := true
			if inputCounter >= 2 {
				success = false
			}
			return core.GoodMap{}, success
		}
		mockApp.writer.Write([]byte("1G"))
		mockApp.writer.Write([]byte("nono"))

		nextState := exchange(mockApp.app, mockApp.exchangeInputParser)
		output := mockApp.writer.String()

		assert.True(t, strings.Contains(output, "Invalid input for selling."))
		assert.False(t, called)
		assert.Equal(t, playerTurn.Name, nextState)
	})

	t.Run("Displays error message on game error", func(t *testing.T) {
		mockApp := newMockApp()
		called := false
		mockApp.game.exchange = func(buy, sell core.GoodMap) error {
			called = true
			return core.GoodsAmountsMismatchError
		}
		mockApp.writer.Write([]byte("1G"))
		mockApp.writer.Write([]byte("2ca"))

		nextState := exchange(mockApp.app, mockApp.exchangeInputParser)
		output := mockApp.writer.String()

		assert.True(t, strings.Contains(output, core.GoodsAmountsMismatchError.Error()))
		assert.True(t, called)
		assert.Equal(t, playerTurn.Name, nextState)
	})
}

func TestParseExchangeInput(t *testing.T) {
	scenarios := []struct {
		in     string
		wanted core.GoodMap
		ok     bool
	}{
		{"2g", core.GoodMap{core.GoodGold: 2}, true},
		{"1G", core.GoodMap{core.GoodGold: 1}, true},
		{"1ca", core.GoodMap{core.GoodCamel: 1}, true},
		{"3cl", core.GoodMap{core.GoodCloth: 3}, true},
		{"asd", nil, false},
		{"4asd", nil, false},
	}

	for _, s := range scenarios {
		t.Run("Test parse exchange input", func(t *testing.T) {
			result, ok := parseExchangeInput(s.in)

			assert.Equal(t, s.ok, ok)
			assert.Equal(t, s.wanted, result)
		})
	}
}
