package core

const GoodsAmountsMismatchError = JaipurError("Goods amounts mismatch")

func (game *Game) Exchange(buy, sell GoodMap) error {
	if err := game.validateExchangeInput(buy, sell); err != nil {
		return err
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

func (game *Game) validateExchangeInput(buy, sell GoodMap) error {
	if game.gameEnded() {
		return GameEndedError
	}

	if game.roundEnded() {
		return RoundEndedError
	}

	amountExchangeSum := Amount(0)
	wantedGoodAmount := Amount(0)
	for good, amount := range buy {
		amountExchangeSum += amount
		if game.cardsOnTable[good] < amount {
			return NotEnoughCardsOnTableError
		}

		switch good {
		case GoodCamel:
		default:
			wantedGoodAmount += amount
		}
	}

	for _, amount := range game.currentPlayer.cards {
		wantedGoodAmount += amount
	}

	if wantedGoodAmount > MaximalPlayerCards {
		return PlayerHasTooManyCardsError
	}

	for good, amount := range sell {
		amountExchangeSum -= amount
		switch good {
		case GoodCamel:
			if game.currentPlayer.herdSize < amount {
				return NotEnoughCardsToSellError
			}
		default:
			if game.currentPlayer.cards[good] < amount {
				return NotEnoughCardsToSellError
			}
		}
	}

	if amountExchangeSum != 0 {
		return GoodsAmountsMismatchError
	}

	return nil
}
