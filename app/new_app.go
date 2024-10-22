package app

import (
	"io"

	"github.com/DAtek/fsm"
)

type AppState = fsm.State[App]

func NewApp(reader io.Reader, writer io.Writer) fsm.IFSM {

	app := &App{
		Reader: reader,
		Writer: writer,
	}

	app.PlayerCommands = PlayerCommandCollection{
		"B": func() fsm.StateName { return buy(app) },
		"S": func() fsm.StateName { return sell(app) },
		"E": func() fsm.StateName { return exchange(app, parseExchangeInput) },
	}

	driver := fsm.NewFSM([]*AppState{
		&gameStarting,
		&playerTurn,
		&roundEnded,
		&gameEnded,
	},
		app,
	)

	return driver
}
