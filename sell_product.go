package jaipur

const NotEnoughCardsError = JaipurError("Player doesn't have enough cards to sell")
const SellingCamelForbiddenError = JaipurError("Selling camels is forbidden")

func (game *game) SellProduct(playerName *Name, productType ProductType, amount Amount) error {
	if productType == ProductCamel {
		return SellingCamelForbiddenError
	}

	player, ok := game.players[*playerName]

	if !ok {
		return PlayerNotExistsError
	}

	if player.cards[productType] < amount {
		return NotEnoughCardsError
	}

	newScore := Score(0)
	for i := Amount(0); i < amount; i++ {
		soldProduct := game.soldProducts
		actualIndex := soldProduct[productType]
		score := coins[productType][actualIndex]
		soldProduct[productType] = actualIndex + 1
		newScore += score
	}

	newScore += getBonus(amount)
	player.score = newScore
	return nil
}
