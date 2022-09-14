package jaipur

const NotEnoughCardsError = JaipurError("Player doesn't have enough cards to sell")
const SellingCamelForbiddenError = JaipurError("Selling camels is forbidden")

func (game *game) SellGoods(playerName *Name, goodType GoodType, amount Amount) error {
	if goodType == GoodCamel {
		return SellingCamelForbiddenError
	}

	player, ok := game.players[*playerName]

	if !ok {
		return PlayerNotExistsError
	}

	if player.cards[goodType] < amount {
		return NotEnoughCardsError
	}

	newScore := Score(0)
	for i := Amount(0); i < amount; i++ {
		soldGoods := game.soldGoods
		actualIndex := soldGoods[goodType]
		score := coins[goodType][actualIndex]
		soldGoods[goodType] = actualIndex + 1
		newScore += score
	}

	newScore += getBonus(amount)
	player.score = newScore
	return nil
}
