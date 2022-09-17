package core

type game struct {
	player1       *player
	player2       *player
	soldGoods     goodMap
	cardsInPack   goodMap
	cardsOnTable  goodMap
	currentPlayer *player
}

func (game *game) RoundEnded() bool {
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
