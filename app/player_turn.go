package app

import (
	"fmt"
	"jaipur/fsm"
)

func doPlayerAction(app *App) fsm.StateName {
	fmt.Fprintln(app.writer, string(app.game.CurrentPlayerName())+", it's your turn")
	playerCards := formatGoodMap(app.game.CurrentPlayerCards())
	cardsOnTable := formatGoodMap(app.game.CardsOnTable())

	fmt.Fprintln(app.writer)
	fmt.Fprintln(app.writer, "Your cards: "+playerCards)
	fmt.Fprintln(app.writer, "Cards on table: "+cardsOnTable)

	action := input(app.reader, app.writer, "Pick an action - (E)xchnge cards | (S)ell cards | (T)ake a card: ")

	switch action {
	case "T":
		return takeCard(app)
	case "E":
		return exchangeCards(app)
	case "S":
		return sellCards(app)
	default:
		fmt.Fprint(app.writer, "Wrong action.\n")
		return playerTurn.Name
	}

}

func takeCard(app *App) fsm.StateName {
	return ""
}

func exchangeCards(app *App) fsm.StateName {
	return ""
}

func sellCards(app *App) fsm.StateName {
	return ""
}
