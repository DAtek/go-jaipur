package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCurrentPlayerCards(t *testing.T) {
	t.Run("Returns current player's cards", func(t *testing.T) {
		game := newGame()

		playerCards := game.CurrentPlayerCards()

		wantedCards := game.currentPlayer.cards.Copy()
		wantedCards[GoodCamel] = game.currentPlayer.herdSize

		assert.Equal(t, wantedCards, playerCards)
	})

	t.Run("Mutating the result won't affect game", func(t *testing.T) {
		game := newGame()

		playerCards := game.CurrentPlayerCards()
		playerCards[GoodDiamond] = 50

		assert.NotEqual(t, game.currentPlayer.cards, playerCards)
	})
}

func TestCardsOnTable(t *testing.T) {
	t.Run("Returns cards on table", func(t *testing.T) {
		game := newGame()

		cards := game.CardsOnTable()

		assert.Equal(t, game.cardsOnTable, cards)
	})

	t.Run("Mutating the result won't affect game", func(t *testing.T) {
		game := newGame()

		cards := game.CardsOnTable()
		cards[GoodDiamond] = 50

		assert.NotEqual(t, game.cardsOnTable, cards)
	})
}

func TestCurrentPlayerName(t *testing.T) {
	game := newGame()

	name := game.CurrentPlayerName()

	assert.Equal(t, game.currentPlayer.name, name)
}

func TestScoreMap(t *testing.T) {
	game := newGame()
	game.player1.score = 24
	game.player2.score = 42

	scores := game.PlayerScores()

	wanted := ScoreMap{
		game.player1.name: game.player1.score,
		game.player2.name: game.player2.score,
	}

	assert.Equal(t, wanted, scores)
}

func TestGameWinner(t *testing.T) {
	player1Name := Name("John")
	player2Name := Name("Jane")
	winnerScenarios := []struct {
		player1Score Score
		player2Score Score
		winner       Name
	}{
		{2, 1, player1Name},
		{0, 2, player2Name},
	}
	for _, s := range winnerScenarios {
		t.Run("Returns winner", func(t *testing.T) {
			game := newGame()
			game.player1.name = player1Name
			game.player2.name = player2Name
			game.player1.sealsOfExcellence = s.player1Score
			game.player2.sealsOfExcellence = s.player2Score
			*game.gameEnded = func() bool { return true }

			winner, err := game.GameWinner()

			assert.Nil(t, err)
			assert.Equal(t, s.winner, winner)
		})
	}

	t.Run("Returns error if game not ended", func(t *testing.T) {
		game := newGame()
		*game.gameEnded = func() bool { return false }

		_, err := game.GameWinner()

		assert.EqualError(t, GameNotEndedError, err.Error())
	})
}
