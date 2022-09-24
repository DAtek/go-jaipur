package app

import (
	"fmt"
	"jaipur/fsm"
)

func sell(app *App) fsm.StateName {
	goodAbbreviation := input(app.reader, app.writer, "Pick a good to sell: ")
	good, ok := goodAbbreviations.find(goodAbbreviation)

	if !ok {
		fmt.Fprint(app.writer, "Invalid input.\n")
		return playerTurn.Name
	}

	err := app.game.Sell(good)

	if err != nil {
		fmt.Fprint(app.writer, err.Error()+"\n\n")
		return playerTurn.Name
	}

	fmt.Fprint(app.writer, clearScreenString)

	if app.game.RoundEnded() {
		return roundEnded.Name
	}

	return playerTurn.Name
}
