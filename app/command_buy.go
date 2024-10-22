package app

import (
	"fmt"

	"github.com/DAtek/fsm"
)

func buy(app *App) fsm.StateName {
	abbreviation := input(app.Reader, app.Writer, "Pick a good to buy: ")
	goodType, ok := goodAbbreviations.find(abbreviation)

	if !ok {
		fmt.Fprint(app.Writer, "Invalid input\n\n")
		return STATE_PLAYER_TURN
	}

	err := app.Game.Buy(goodType)
	if err != nil {
		fmt.Fprintf(app.Writer, "%s\n\n", err)
		return STATE_PLAYER_TURN
	}

	fmt.Fprint(app.Writer, clearScreenString)

	if app.Game.RoundEnded() {
		return STATE_ROUND_ENDED
	}

	return STATE_PLAYER_TURN
}
