package app

import (
	"bytes"
	"io"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStartGame(t *testing.T) {
	t.Run("Returns original state when player name is empty", func(t *testing.T) {
		m := newMockApp()
		m.app.Reader = &bytes.Buffer{}

		nextState := gameStarting.Transit(m.app)

		assert.Equal(t, STATE_GAME_STARTING, nextState)
	})

	t.Run("Collects player names", func(t *testing.T) {
		m := newMockApp()
		index := 0
		playerIdx := 0
		players := [][]byte{[]byte("player1"), []byte("player2")}
		reader := &mockReader{
			Read_: func(p []byte) (n int, err error) {
				if index >= len(players[playerIdx]) {
					playerIdx++
					index = 0
					return 0, io.EOF
				}

				n = copy(p, players[playerIdx][index:])
				index += n

				return n, nil
			},
		}
		m.app.Reader = reader

		nextState := gameStarting.Transit(m.app)
		output := m.writer.String()

		assert.Equal(t, STATE_PLAYER_TURN, nextState)
		assert.True(t, strings.Contains(output, clearScreenString))
	})

	t.Run("Asks for player names", func(t *testing.T) {
		m := newMockApp()
		index := 0
		playerIdx := 0
		players := [][]byte{[]byte("player1"), []byte("player2")}
		reader := &mockReader{
			Read_: func(p []byte) (n int, err error) {
				if index >= len(players[playerIdx]) {
					playerIdx++
					index = 0
					return 0, io.EOF
				}

				n = copy(p, players[playerIdx][index:])
				index += n

				return n, nil
			},
		}
		m.app.Reader = reader

		gameStarting.Transit(m.app)

		output := m.writer.String()

		assert.True(t, strings.Contains(output, "Enter player 1 name: "))
		assert.True(t, strings.Contains(output, "Enter player 2 name: "))
	})
}
