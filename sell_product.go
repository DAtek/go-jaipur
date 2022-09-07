package jaipur

const NotEnoughCardsError = JaipurError("Player doesn't have enough cards to sell")

func (game *game) SellProduct(playerName *Name, goodsType ProductType, amount Amount) error {
	newScore := Score(0)
	players := game.players
	player, ok := players[*playerName]

	if !ok {
		return PlayerNotExistsError
	}

	if player.cards[goodsType] < amount {
		return NotEnoughCardsError
	}

	for i := Amount(0); i < amount; i++ {
		soldProduct := game.soldProducts
		actualIndex := soldProduct[goodsType]
		score := coins[goodsType][actualIndex]
		soldProduct[goodsType] = actualIndex + 1
		newScore += score
	}

	newScore += getBonus(amount)
	player.score = newScore
	return nil
}
