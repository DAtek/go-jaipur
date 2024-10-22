package app

import (
	"fmt"
	"jaipur/core"
	"strconv"
	"strings"

	"github.com/DAtek/fsm"
)

type inputParser func(string) (core.GoodMap, bool)

func exchange(app *App, parseInput inputParser) fsm.StateName {
	buyString := input(app.Reader, app.Writer, "Buy goods eg. 2G, 1Si: ")
	buy, ok := parseInput(buyString)
	if !ok {
		fmt.Fprint(app.Writer, "Invalid input for buying.\n\n")
		return STATE_PLAYER_TURN
	}

	sellString := input(app.Reader, app.Writer, "Sell goods eg. 2G, 1Ca: ")
	sell, ok := parseInput(sellString)

	if !ok {
		fmt.Fprint(app.Writer, "Invalid input for selling.\n\n")
		return STATE_PLAYER_TURN
	}

	err := app.Game.Exchange(buy, sell)

	if err != nil {
		fmt.Fprint(app.Writer, err.Error()+"\n\n")
		return STATE_PLAYER_TURN
	}

	fmt.Fprint(app.Writer, clearScreenString)
	return STATE_PLAYER_TURN
}

func parseExchangeInput(input string) (core.GoodMap, bool) {
	parts := strings.Split(input, ", ")
	goodMap := core.GoodMap{}

	for _, item := range parts {
		goodType, amount, ok := parseGoodMapItem(item)

		if !ok {
			return nil, false
		}

		goodMap[goodType] = amount
	}

	return goodMap, true
}

func parseGoodMapItem(item string) (core.GoodType, core.Amount, bool) {
	amountInt, err := strconv.Atoi(item[0:1])

	if err != nil {
		return 0, 0, false
	}

	amount := core.Amount(amountInt)
	goodType, ok := goodAbbreviations.find(item[1:])

	if !ok {
		return 0, 0, false
	}

	return goodType, amount, true
}
