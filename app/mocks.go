package app

import (
	"bytes"
	"jaipur/core"
	"jaipur/fsm"
)

type mockGame struct {
	currentPlayerName  core.Name
	takeCard           func(card core.GoodType) error
	sellGoods          func(card core.GoodType) error
	currentPlayerCards core.GoodMap
	cardsOnTable       core.GoodMap
	roundEnded         bool
}

type mockApp struct {
	reader *bytes.Buffer
	writer *bytes.Buffer
	app    *App
	game   *mockGame
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

func (game *mockGame) RoundEnded() bool {
	return game.roundEnded
}

func (game *mockGame) SellGoods(good core.GoodType) error {
	return game.sellGoods(good)
}

func newMockApp() *mockApp {
	reader := &bytes.Buffer{}
	writer := &bytes.Buffer{}
	game := &mockGame{
		currentPlayerName: "Max",
		roundEnded:        false,
		takeCard:          func(card core.GoodType) error { return nil },
		sellGoods:         func(card core.GoodType) error { return nil },
	}

	return &mockApp{
		reader: reader,
		writer: writer,
		game:   game,
		app: &App{
			reader: reader,
			writer: writer,
			game:   game,
		},
	}
}

func newMockPlayerCommand() *playerCommandCollection {
	return &playerCommandCollection{
		TakeCard:      func() fsm.StateName { return "" },
		ExchangeCards: func() fsm.StateName { return "" },
		SellCards:     func() fsm.StateName { return "" },
	}
}
