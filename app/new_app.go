package app

import (
	"bufio"
	"io"
	"jaipur/fsm"
)

func NewApp(reader io.Reader, writer io.Writer) *App {
	driver := fsm.FSM{States: []*fsm.State{&gameStart, &playerTurn, &roundEnded, &gameEnded}}

	app := &App{
		fsm:    &driver,
		reader: bufio.NewReader(reader),
		writer: writer,
	}

	playerCommands := &playerCommandCollection{}

	startTransition := func() fsm.StateName {
		return askForNames(app)
	}
	gameStart.Transit = &startTransition

	playerTurnTransition := func() fsm.StateName {
		return doPlayerAction(app, playerCommands)
	}
	playerTurn.Transit = &playerTurnTransition

	playerBuyTransition := func() fsm.StateName {
		return buy(app)
	}
	playerCommands.Buy = playerBuyTransition

	playerSellTransition := func() fsm.StateName {
		return sell(app)
	}
	playerCommands.Sell = playerSellTransition

	playerExchangeTransition := func() fsm.StateName {
		return exchange(app, parseExchangeInput)
	}
	playerCommands.Exchange = playerExchangeTransition

	return app
}
