package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoundEnded(t *testing.T) {
	tokenDepletedScenarios := []GoodMap{
		{
			GoodDiamond: 5,
			GoodGold:    5,
			GoodCloth:   7,
		},
		{
			GoodLeather: 9,
			GoodSilver:  5,
			GoodSpice:   7,
		},
	}

	for _, s := range tokenDepletedScenarios {
		t.Run("Round ended when 3 tokens depleted", func(t *testing.T) {
			game := newGameMock()
			game.soldGoods = s

			assert.Equal(t, true, roundEnded(game))
		})
	}

	t.Run("Round ended when there are less than 5 cards on table", func(t *testing.T) {
		game := newGameMock()
		game.cardsOnTable = GoodMap{GoodCloth: 4}

		assert.Equal(t, true, roundEnded(game))
	})

}
