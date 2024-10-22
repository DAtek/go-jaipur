package app

import (
	"fmt"
	"jaipur/core"

	"github.com/DAtek/fsm"
)

var gameStarting = AppState{
	Name: STATE_GAME_STARTING,
	Transit: func(ctx *App) fsm.StateName {
		name := input(ctx.Reader, ctx.Writer, "Enter player 1 name: ")
		player1 := core.Name(name)
		name = input(ctx.Reader, ctx.Writer, "Enter player 2 name: ")
		player2 := core.Name(name)

		e := error(nil)
		ctx.Game, e = core.NewGame(player1, player2)
		switch e {
		case nil:
			fmt.Fprint(ctx.Writer, clearScreenString)
			return STATE_PLAYER_TURN
		default:
			fmt.Fprintf(ctx.Writer, "\nError: %s\n\n", e.Error())
			return STATE_GAME_STARTING
		}
	},
}
