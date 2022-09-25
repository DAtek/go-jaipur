package app

import (
	"fmt"
	"io"
	"jaipur/core"
	"jaipur/fsm"
)

func startGame(app *App, input func(io.Reader, io.Writer, string) string) fsm.StateName {
	name := input(app.reader, app.writer, "Enter player 1 name: ")
	player1 := core.Name(name)
	name = input(app.reader, app.writer, "Enter player 2 name: ")
	player2 := core.Name(name)

	e := error(nil)
	app.game, e = core.NewGame(player1, player2)
	switch e {
	case nil:
		fmt.Fprint(app.writer, clearScreenString)
		return playerTurn.Name
	default:
		fmt.Fprintf(app.writer, "\nError: %s\n\n", e.Error())
		return gameStart.Name
	}
}
