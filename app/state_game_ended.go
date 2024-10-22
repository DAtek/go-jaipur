package app

import (
	"fmt"

	"github.com/DAtek/fsm"
)

var gameEnded = AppState{
	Name: STATE_GAME_ENDED,
	Transit: func(ctx *App) fsm.StateName {
		winner, err := ctx.Game.GameWinner()

		if err != nil {
			panic("Unexpected error: " + err.Error())
		}

		fmt.Fprint(ctx.Writer, "Congratulations, "+string(winner)+"! You won!\n")
		input(ctx.Reader, ctx.Writer, "Press enter to quit")
		return fsm.STATE_FINAL
	},
}
