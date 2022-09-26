package core

func (game *Game) Buy(card GoodType) error {
	if game.gameEnded() {
		return GameEndedError
	}

	if game.roundEnded() {
		return RoundEndedError
	}

	sumPlayerCards := Amount(0)
	for _, amount := range game.currentPlayer.cards {
		sumPlayerCards += amount
	}

	if sumPlayerCards == MaximumCardsInHand {
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

	game.changeCurrentPlayer()

	return nil
}
