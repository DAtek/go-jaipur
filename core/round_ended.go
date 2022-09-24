package core

func (game *Game) RoundEnded() bool {
	return (*game.roundEnded)()
}

func roundEnded(game *Game) bool {
	depletedTokens := Amount(0)
	for good, amount := range game.soldGoods {
		if Amount(len(coins[good])) == amount {
			depletedTokens++
		}
	}

	if depletedTokens >= 3 {
		return true
	}

	cards := Amount(0)
	for _, amount := range game.cardsOnTable {
		cards += amount
	}

	return cards < 5
}
