package app

import (
	"fmt"
	"jaipur/fsm"
	"strings"
)

func doPlayerAction(app *App, playerCommand *playerCommandCollection) fsm.StateName {
	fmt.Fprintln(app.writer, string(app.game.CurrentPlayerName())+", it's your turn")
	playerCards := formatGoodMap(app.game.CurrentPlayerCards())
	cardsOnTable := formatGoodMap(app.game.CardsOnTable())

	fmt.Fprintln(app.writer)
	fmt.Fprintln(app.writer, "Your cards: "+playerCards)
	fmt.Fprintln(app.writer, "Cards on table: "+cardsOnTable)

	action := input(app.reader, app.writer, "Pick an action - (E)xchnge cards | (S)ell cards | (T)ake a card: ")
	action = strings.ToUpper(action)

	switch action {
	case "T":
		return playerCommand.TakeCard()
	case "E":
		return playerCommand.ExchangeCards()
	case "S":
		return playerCommand.SellCards()
	default:
		fmt.Fprint(app.writer, "Wrong action.\n\n")
		return playerTurn.Name
	}
}
