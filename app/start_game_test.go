package app

import (
	"bytes"
	"fmt"
	"io"
	"jaipur/core"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartGame(t *testing.T) {
	t.Run("Returns original state when players have the same name", func(t *testing.T) {
		m := newMockApp()
		input := func(r io.Reader, w io.Writer, s string) string {
			return "a"
		}

		nextState := startGame(m.app, input)
		output := m.writer.String()

		assert.Equal(t, gameStart.Name, nextState)
		assert.True(t, strings.Contains(output, core.SameNamesError.Error()))
	})

	t.Run("Collects player names", func(t *testing.T) {
		m := newMockApp()
		players := []string{"player1", "player2"}
		index := -1
		input := func(r io.Reader, w io.Writer, s string) string {
			index += 1
			return players[index]
		}

		nextState := startGame(m.app, input)
		output := m.writer.String()

		assert.Equal(t, playerTurn.Name, nextState)
		assert.True(t, strings.Contains(output, clearScreenString))
	})

	t.Run("Asks for player names", func(t *testing.T) {
		m := newMockApp()
		buf := bytes.Buffer{}
		m.writer = &buf
		input := func(r io.Reader, w io.Writer, s string) string {
			buf.WriteString(s)
			return ""
		}

		startGame(m.app, input)

		output := m.writer.String()
		fmt.Printf("output: %v\n", output)
		assert.True(t, strings.Contains(output, "Enter player 1 name: "))
		assert.True(t, strings.Contains(output, "Enter player 2 name: "))
	})
}
