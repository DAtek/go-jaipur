package app

import (
	"bytes"
	"jaipur/core"
)

type mockGame struct {
	currentPlayerName  core.Name
	buy                func(card core.GoodType) error
	sell               func(card core.GoodType) error
	exchange           func(buy core.GoodMap, sell core.GoodMap) error
	currentPlayerCards core.GoodMap
	cardsOnTable       core.GoodMap
	roundEnded         bool
	gameEnded          bool
	roundWinner        func() (core.Name, error)
	playerScores       core.ScoreMap
	finishRound        func() error
	gameWinner         func() (core.Name, error)
}

type mockApp struct {
	reader              *bytes.Buffer
	writer              *bytes.Buffer
	app                 *App
	game                *mockGame
	exchangeInputParser inputParser
}

func (game *mockGame) CurrentPlayerName() core.Name {
	return game.currentPlayerName
}

func (game *mockGame) CurrentPlayerCards() core.GoodMap {
	return game.currentPlayerCards
}

func (game *mockGame) CardsOnTable() core.GoodMap {
	return game.cardsOnTable
}

func (game *mockGame) Buy(card core.GoodType) error {
	return game.buy(card)
}

func (game *mockGame) RoundEnded() bool {
	return game.roundEnded
}

func (game *mockGame) GameEnded() bool {
	return game.gameEnded
}

func (game *mockGame) Sell(good core.GoodType) error {
	return game.sell(good)
}

func (game *mockGame) Exchange(buy core.GoodMap, sell core.GoodMap) error {
	return game.exchange(buy, sell)
}

func (game *mockGame) RoundWinner() (core.Name, error) {
	return game.roundWinner()
}

func (game *mockGame) PlayerScores() core.ScoreMap {
	return game.playerScores
}

func (game *mockGame) FinishRound() error {
	return game.finishRound()
}

func (game *mockGame) GameWinner() (core.Name, error) {
	return game.gameWinner()
}

var _ core.IGame = &mockGame{}

func newMockApp() *mockApp {
	game := &mockGame{
		currentPlayerName: "Max",
		roundEnded:        false,
		buy:               func(card core.GoodType) error { return nil },
		sell:              func(card core.GoodType) error { return nil },
		exchange:          func(buy, sell core.GoodMap) error { return nil },
	}

	roundWinner := func() (core.Name, error) {
		return "", nil
	}
	gameWinner := func() (core.Name, error) {
		return "", nil
	}
	game.roundWinner = roundWinner
	game.gameWinner = gameWinner
	game.finishRound = func() error { return nil }
	reader := &bytes.Buffer{}
	writer := &bytes.Buffer{}

	return &mockApp{
		reader: reader,
		writer: writer,
		game:   game,
		app: &App{
			Reader:         reader,
			Writer:         writer,
			Game:           game,
			PlayerCommands: PlayerCommandCollection{},
		},
		exchangeInputParser: func(s string) (core.GoodMap, bool) {
			return core.GoodMap{}, true
		},
	}
}

type mockReader struct {
	Read_ func(p []byte) (n int, err error)
}

func (m *mockReader) Read(p []byte) (n int, err error) { return m.Read_(p) }
