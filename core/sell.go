package core

const SellingCamelForbiddenError = JaipurError("Selling camels is forbidden")
const GoodSoldOutError = JaipurError("This good is sold out")

func (game *game) Sell(goodType GoodType) error {
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

	soldGoods := game.soldGoods
	actualIndex := soldGoods[goodType]
	totalCoins := Amount(len(coins[goodType]))

	if actualIndex == totalCoins {
		return GoodSoldOutError
	}

	for i := Amount(0); i < amount; i++ {
		actualIndex = soldGoods[goodType]
		if actualIndex == totalCoins {
			return nil
		}
		score := coins[goodType][actualIndex]
		soldGoods[goodType] = actualIndex + 1
		game.currentPlayer.score += score
		game.currentPlayer.cards[goodType]--
	}

	game.currentPlayer.score += getBonus(amount)
	game.changeCurrentPlayer()
	return nil
}
