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

	startTransition := func() fsm.StateName {
		return startGame(app)
	}
	gameStart.Transit = &startTransition

	playerTurnTransition := func() fsm.StateName {
		return doPlayerAction(app, playerCommands)
	}
	playerTurn.Transit = &playerTurnTransition

	finishRoundTransition := func() fsm.StateName {
		return finishRound(app)
	}
	roundEnded.Transit = &finishRoundTransition

	endGameTransition := func() fsm.StateName {
		return endGame(app)
	}
	gameEnded.Transit = &endGameTransition

	playerBuy := func() fsm.StateName {
		return buy(app)
	}
	playerCommands.Buy = &playerBuy

	playerSell := func() fsm.StateName {
		return sell(app)
	}
	playerCommands.Sell = &playerSell

	playerExchange := func() fsm.StateName {
		return exchange(app, parseExchangeInput)
	}
	playerCommands.Exchange = &playerExchange

	return app
}
