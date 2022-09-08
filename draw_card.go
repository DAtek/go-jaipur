package jaipur

const NotEnoughCardsOnTableError = JaipurError("Not enough cards on the table")
const PlayerHasTooManyCardsError = JaipurError("PlayerHasTooManyCards")

func (game *game) DrawCard(playerName Name, card ProductType) error {
	player, ok := game.players[playerName]

	if !ok {
		return PlayerNotExistsError
	}

	sumPlayerCards := Amount(0)
	for _, amount := range player.cards {
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
	case ProductCamel:
		player.herdSize += amount
		game.cardsOnTable[card] -= amount
		game.moveCardsFromPackToTable(amount)
	default:
		player.cards[card] += 1
		game.cardsOnTable[card] -= 1
		game.moveCardsFromPackToTable(1)
	}

	return nil
}
