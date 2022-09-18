package core

const GoodsAmountsMismatchError = JaipurError("Goods amounts mismatch")

func (game *Game) ExchangeGoods(buy GoodMap, sell GoodMap) error {
	if game.gameEnded() {
		return GameEndedError
	}

	if game.roundEnded() {
		return RoundEndedError
	}

	amountExchangeSum := Amount(0)

	for key, value := range buy {
		amountExchangeSum += value
		if game.cardsOnTable[key] < value {
			return NotEnoughCardsOnTableError
		}
	}

	for key, value := range sell {
		amountExchangeSum -= value
		if game.currentPlayer.cards[key] < value {
			return NotEnoughCardsToSellError
		}
	}

	if amountExchangeSum != 0 {
		return GoodsAmountsMismatchError
	}

	for key, value := range sell {
		game.cardsOnTable[key] += value
		game.currentPlayer.cards[key] -= value
	}

	for key, value := range buy {
		game.currentPlayer.cards[key] += value
		game.cardsOnTable[key] -= value
	}

	game.changeCurrentPlayer()

	return nil
}
