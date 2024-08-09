package app

import (
	"fmt"
	"jaipur/fsm"
	"strings"
)

func doPlayerAction(app *App, playerCommands playerCommandCollection) fsm.StateName {
	fmt.Fprintln(app.writer, string(app.game.CurrentPlayerName())+", it's your turn")
	cardsOnTable := formatGoodMap(app.game.CardsOnTable())
	playerCards := formatGoodMap(app.game.CurrentPlayerCards())

	fmt.Fprintln(app.writer)
	fmt.Fprintln(app.writer, "Your cards: "+playerCards)
	fmt.Fprintln(app.writer, "Cards on table: "+cardsOnTable+"\n")

	action := input(app.reader, app.writer, "Pick an action - (E)xchange | (S)ell | (B)uy: ")
	action = strings.ToUpper(action)
	command, ok := playerCommands[action]

	if !ok {
		fmt.Fprint(app.writer, "Wrong action.\n\n")
		return playerTurn.Name
	}

	return command()
}
