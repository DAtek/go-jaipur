package app

import (
	"fmt"
	"jaipur/core"
	"jaipur/fsm"
	"strconv"
)

type inputParser func(string) (core.GoodMap, bool)

func exchange(app *App, parseInput inputParser) fsm.StateName {
	buyString := input(app.reader, app.writer, "Buy goods eg. 2G, 1S: ")
	buy, ok := parseInput(buyString)

	if !ok {
		fmt.Fprint(app.writer, "Invalid input for buying.\n\n")
		return playerTurn.Name
	}

	sellString := input(app.reader, app.writer, "Sell goods eg. 2G, 1Ca: ")
	sell, ok := parseInput(sellString)

	if !ok {
		fmt.Fprint(app.writer, "Invalid input for selling.\n\n")
		return playerTurn.Name
	}

	err := app.game.Exchange(buy, sell)

	if err != nil {
		fmt.Fprint(app.writer, err.Error()+"\n\n")
		return playerTurn.Name
	}

	fmt.Fprint(app.writer, clearScreenString)
	return playerTurn.Name
}

func parseExchangeInput(input string) (core.GoodMap, bool) {
	amountInt, err := strconv.Atoi(input[0:1])
	if err != nil {
		return nil, false
	}
	amount := core.Amount(amountInt)
	goodType, ok := goodAbbreviations.find(input[1:])
	if !ok {
		return nil, false
	}
	return core.GoodMap{goodType: amount}, true
}
