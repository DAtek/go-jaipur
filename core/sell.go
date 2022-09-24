package core

const SellingCamelForbiddenError = JaipurError("Selling camels is forbidden")

func (game *Game) Sell(goodType GoodType) error {
	if game.gameEnded() {
		return GameEndedError
	}

	if game.roundEnded() {
		return RoundEndedError
	}

	if goodType == GoodCamel {
		return SellingCamelForbiddenError
	}

	amount := game.currentPlayer.cards[goodType]
	if amount == 0 {
		return NotEnoughCardsToSellError
	}

	if amount == 1 {
		for _, expensiveGood := range expensiveGoods {
			if goodType == expensiveGood {
				return NotEnoughCardsToSellError
			}
		}
	}

	for i := Amount(0); i < amount; i++ {
		soldGoods := game.soldGoods
		actualIndex := soldGoods[goodType]
		score := coins[goodType][actualIndex]
		soldGoods[goodType] = actualIndex + 1
		game.currentPlayer.score += score
	}

	game.currentPlayer.score += getBonus(amount)
	game.currentPlayer.cards[goodType] -= amount
	game.changeCurrentPlayer()
	return nil
}
