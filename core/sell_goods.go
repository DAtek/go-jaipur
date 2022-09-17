package core

const SellingCamelForbiddenError = JaipurError("Selling camels is forbidden")

func (game *game) SellGoods(goodType GoodType) error {
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

	newScore := Score(0)
	for i := Amount(0); i < amount; i++ {
		soldGoods := game.soldGoods
		actualIndex := soldGoods[goodType]
		score := coins[goodType][actualIndex]
		soldGoods[goodType] = actualIndex + 1
		newScore += score
	}

	newScore += getBonus(amount)
	game.currentPlayer.cards[goodType] -= amount
	game.currentPlayer.score = newScore
	game.changeCurrentPlayer()
	return nil
}
