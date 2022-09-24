package app

import (
	"fmt"
	"jaipur/fsm"
)

func endGame(app *App) fsm.StateName {
	winner, err := app.game.GameWinner()

	if err != nil {
		panic("Unexpected error: " + err.Error())
	}

	fmt.Fprint(app.writer, "Congratulations, "+string(winner)+"! You won!\n")
	input(app.reader, app.writer, "Press enter to quit")
	return finalState.Name
}
