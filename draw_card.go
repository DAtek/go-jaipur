package jaipur

func (game *game) DrawCard(playerName Name, card ProductType, amount Amount) error {
	player, ok := game.players[playerName]

	if !ok {
		return PlayerNotExistsError
	}

	player.cards[card] += amount
	game.cardsOnTable[card] -= amount
	game.moveCardsFromPackToTable(amount)
	return nil
}
