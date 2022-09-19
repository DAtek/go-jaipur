package app

import (
	"fmt"
	"jaipur/core"
	"jaipur/fsm"
)

func askForNames(app *App) fsm.StateName {
	name := input(app.reader, app.writer, "Enter player 1 name: ")
	player1 := core.Name(name)
	name = input(app.reader, app.writer, "Enter player 2 name: ")
	player2 := core.Name(name)

	e := error(nil)
	app.game, e = core.NewGame(player1, player2)

	switch e {
	case core.SameNamesError:
		fmt.Fprintf(app.writer, "\nError: %s\n\n", core.SameNamesError.Error())
		return gameStart.Name
	default:
		fmt.Fprintln(app.writer)
		return playerTurn.Name
	}
}
