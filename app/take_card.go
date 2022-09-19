package app

import (
	"fmt"
	"jaipur/fsm"
)

func takeCard(app *App) fsm.StateName {
	abbreviation := input(app.reader, app.writer, "Take a card: ")
	goodType, ok := goodAbbreviations.find(abbreviation)
	fmt.Fprintln(app.writer)

	if !ok {
		fmt.Fprint(app.writer, "Invalid input\n\n")
		return playerTurn.Name
	}

	err := app.game.TakeCard(goodType)
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
