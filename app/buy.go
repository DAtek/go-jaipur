package app

import (
	"fmt"
	"jaipur/fsm"
)

func buy(app *App) fsm.StateName {
	abbreviation := input(app.reader, app.writer, "Pick a good to buy: ")
	fmt.Fprint(app.writer, "picked good: '"+abbreviation+"'")
	goodType, ok := goodAbbreviations.find(abbreviation)

	if !ok {
		fmt.Fprint(app.writer, "Invalid input\n\n")
		return playerTurn.Name
	}

	err := app.game.Buy(goodType)
	if err != nil {
		fmt.Fprintf(app.writer, "%s\n\n", err)
		return playerTurn.Name
	}

	fmt.Fprint(app.writer, clearScreenString)

	if app.game.RoundEnded() {
		return roundEnded.Name
	}

	return playerTurn.Name
}
