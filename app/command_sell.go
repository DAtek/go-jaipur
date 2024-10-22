package app

import (
	"fmt"

	"github.com/DAtek/fsm"
)

func sell(app *App) fsm.StateName {
	goodAbbreviation := input(app.Reader, app.Writer, "Pick a good to sell: ")
	good, ok := goodAbbreviations.find(goodAbbreviation)

	if !ok {
		fmt.Fprint(app.Writer, "Invalid input.\n")
		return STATE_PLAYER_TURN
	}

	err := app.Game.Sell(good)

	if err != nil {
		fmt.Fprint(app.Writer, err.Error()+"\n\n")
		return STATE_PLAYER_TURN
	}

	fmt.Fprint(app.Writer, clearScreenString)

	if app.Game.RoundEnded() {
		return STATE_ROUND_ENDED
	}

	return STATE_PLAYER_TURN
}
