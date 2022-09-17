package core

const PlayerHasTooManyCardsError = JaipurError("PlayerHasTooManyCards")

func (game *game) TakeCard(card GoodType) error {
	sumPlayerCards := Amount(0)
	for _, amount := range game.currentPlayer.cards {
		sumPlayerCards += amount
	}

	if sumPlayerCards == 7 {
		return PlayerHasTooManyCardsError
	}

	amount, ok := game.cardsOnTable[card]
	if !ok || amount < 1 {
		return NotEnoughCardsOnTableError
	}

	switch card {
	case GoodCamel:
		game.currentPlayer.herdSize += amount
		game.cardsOnTable[card] -= amount
		game.moveCardsFromPackToTable(amount)
	default:
		game.currentPlayer.cards[card] += 1
		game.cardsOnTable[card] -= 1
		game.moveCardsFromPackToTable(1)
	}

	return nil
}
