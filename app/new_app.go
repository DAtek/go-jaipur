package app

import (
	"io"
	"jaipur/fsm"
)

func NewApp(reader io.Reader, writer io.Writer) *App {
	driver := fsm.NewFSM([]*fsm.State{&gameStart, &playerTurn, &roundEnded, &gameEnded, &finalState})

	app := &App{
		fsm:    driver,
		reader: reader,
		writer: writer,
	}

	playerCommands := &playerCommandCollection{
		Buy:      func() fsm.StateName { return buy(app) },
		Sell:     func() fsm.StateName { return sell(app) },
		Exchange: func() fsm.StateName { return exchange(app, parseExchangeInput) },
	}

	gameStart.Transit = func() fsm.StateName { return startGame(app, input) }
	playerTurn.Transit = func() fsm.StateName { return doPlayerAction(app, playerCommands) }
	roundEnded.Transit = func() fsm.StateName { return finishRound(app) }
	gameEnded.Transit = func() fsm.StateName { return endGame(app) }
	return app
}
