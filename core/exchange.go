package core

const GoodsAmountsMismatchError = JaipurError("Goods amounts mismatch")

func (game *game) Exchange(buy, sell GoodMap) error {
	if err := game.validateExchangeInput(buy, sell); err != nil {
		return err
	}

	for card, amount := range sell {
		game.cardsOnTable[card] += amount
		switch card {
		case GoodCamel:
			game.currentPlayer.herdSize -= amount
		default:
			game.currentPlayer.cards[card] -= amount
		}
	}

	for card, amount := range buy {
		switch card {
		case GoodCamel:
			game.currentPlayer.herdSize += amount
		default:
			game.currentPlayer.cards[card] += amount
		}
		game.cardsOnTable[card] -= amount
	}

	game.changeCurrentPlayer()

	return nil
}

func (game *game) validateExchangeInput(buy, sell GoodMap) error {
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

	for good, amount := range sell {
		amountExchangeSum -= amount
		switch good {
		case GoodCamel:
			if game.currentPlayer.herdSize < amount {
				return NotEnoughCardsToSellError
			}
		default:
			wantedGoodAmount -= amount
			if game.currentPlayer.cards[good] < amount {
				return NotEnoughCardsToSellError
			}
		}
	}

	if wantedGoodAmount > MaximumCardsInHand {
		return PlayerHasTooManyCardsError
	}

	if amountExchangeSum != 0 {
		return GoodsAmountsMismatchError
	}

	return nil
}
