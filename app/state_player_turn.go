package app

import (
	"fmt"
	"strings"

	"github.com/DAtek/fsm"
)

var playerTurn = AppState{
	Name: STATE_PLAYER_TURN,
	Transit: func(ctx *App) fsm.StateName {
		fmt.Fprintln(ctx.Writer, string(ctx.Game.CurrentPlayerName())+", it's your turn")
		cardsOnTable := formatGoodMap(ctx.Game.CardsOnTable())
		playerCards := formatGoodMap(ctx.Game.CurrentPlayerCards())

		fmt.Fprintln(ctx.Writer)
		fmt.Fprintln(ctx.Writer, "Your cards: "+playerCards)
		fmt.Fprintln(ctx.Writer, "Cards on table: "+cardsOnTable+"\n")

		action := input(ctx.Reader, ctx.Writer, "Pick an action - (E)xchange | (S)ell | (B)uy: ")
		action = strings.ToUpper(action)
		command, ok := ctx.PlayerCommands[action]

		if !ok {
			fmt.Fprint(ctx.Writer, "Wrong action.\n\n")
			return STATE_PLAYER_TURN
		}

		return command()
	},
}
