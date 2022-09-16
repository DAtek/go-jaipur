package core

const GoodsAmountsMismatchError = JaipurError("Goods amounts mismatch")

func (game *game) ExchangeGoods(playerName Name, buy goodMap, sell goodMap) error {
	player, ok := game.players[playerName]

	if !ok {
		return PlayerNotExistsError
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
		if player.cards[key] < value {
			return NotEnoughCardsToSellError
		}
	}

	if amountExchangeSum != 0 {
		return GoodsAmountsMismatchError
	}

	for key, value := range sell {
		game.cardsOnTable[key] += value
		player.cards[key] -= value
	}

	for key, value := range buy {
		player.cards[key] += value
		game.cardsOnTable[key] -= value
	}

	return nil
}
