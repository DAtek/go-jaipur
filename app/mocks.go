package app

import (
	"bytes"
	"jaipur/core"
)

type mockApp struct {
	reader *bytes.Buffer
	writer *bytes.Buffer
	app    *App
}

type mockGame struct {
	currentPlayerName  core.Name
	takeCard           func(card core.GoodType) error
	currentPlayerCards core.GoodMap
	cardsOnTable       core.GoodMap
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

func (game *mockGame) TakeCard(card core.GoodType) error {
	return game.takeCard(card)
}

func newMockApp() *mockApp {
	reader := &bytes.Buffer{}
	writer := &bytes.Buffer{}

	return &mockApp{
		reader: reader,
		writer: writer,
		app: &App{
			reader: reader,
			writer: writer,
			game: &mockGame{
				currentPlayerName: "Max",
			},
		},
	}
}
