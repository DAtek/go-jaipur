package app

import (
	"bufio"
	"io"
	"jaipur/fsm"
)

func NewApp(reader io.Reader, writer io.Writer) *App {
	driver := fsm.FSM{States: []*fsm.State{&gameStart, &playerTurn, &roundEnded, &gameEnded, &finalState}}

	app := &App{
		fsm:    &driver,
		reader: bufio.NewReader(reader),
		writer: writer,
	}

	playerCommands := &playerCommandCollection{}
	gameStart.Transit = func() fsm.StateName { return startGame(app) }
	playerTurn.Transit = func() fsm.StateName { return doPlayerAction(app, playerCommands) }
	roundEnded.Transit = func() fsm.StateName { return finishRound(app) }
	gameEnded.Transit = func() fsm.StateName { return endGame(app) }
	playerCommands.Buy = func() fsm.StateName { return buy(app) }
	playerCommands.Sell = func() fsm.StateName { return sell(app) }
	playerCommands.Exchange = func() fsm.StateName { return exchange(app, parseExchangeInput) }
	return app
}
