package app

import (
	"fmt"

	"github.com/DAtek/fsm"
)

var roundEnded = AppState{
	Name: STATE_ROUND_ENDED,
	Transit: func(ctx *App) fsm.StateName {
		winner, _ := ctx.Game.RoundWinner()
		fmt.Fprint(ctx.Writer, "Winner of the round: "+string(winner)+"\n")
		playerScores := ctx.Game.PlayerScores()
		for name, score := range playerScores {
			fmt.Fprintf(ctx.Writer, "%s's score: %d\n", name, score)
		}
		input(ctx.Reader, ctx.Writer, "Press enter to continue")
		err := ctx.Game.FinishRound()

		if err != nil {
			panic("Unexpected error: " + err.Error())
		}

		fmt.Fprint(ctx.Writer, clearScreenString)

		if ctx.Game.GameEnded() {
			return STATE_GAME_ENDED
		}

		return STATE_PLAYER_TURN
	},
}
