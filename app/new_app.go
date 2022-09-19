package app

import (
	"io"
	"jaipur/fsm"
)

func NewApp(reader io.Reader, writer io.Writer) *App {
	driver := fsm.FSM{States: []*fsm.State{&gameStart, &playerTurn, &roundEnded, &gameEnded}}

	app := &App{
		fsm:    &driver,
		reader: reader,
		writer: writer,
	}

	playerCommands := &playerCommandCollection{}

	gameStart.Transit = func() fsm.StateName {
		return askForNames(app)
	}

	playerTurn.Transit = func() fsm.StateName {
		return doPlayerAction(app, playerCommands)
	}

	playerCommands.TakeCard = func() fsm.StateName {
		return takeCard(app)
	}

	return app
}
