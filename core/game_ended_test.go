package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGameEnded(t *testing.T) {
	scenarios := []struct {
		player1SealsOfExcellence Score
		player2SealsOfExcellence Score
		wanted                   bool
	}{
		{0, 0, false},
		{1, 0, false},
		{2, 1, true},
		{2, 0, true},
	}

	for _, s := range scenarios {
		t.Run("Game ended", func(t *testing.T) {
			game := newGameMock()
			game.player1.sealsOfExcellence = s.player1SealsOfExcellence
			game.player2.sealsOfExcellence = s.player2SealsOfExcellence

			assert.Equal(t, s.wanted, gameEnded(game))
		})
	}
}
