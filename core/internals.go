package core

import (
	"math/rand"
)

func (game *game) moveCardsFromPackToTable(amount Amount) {
	allCards := game.getAllCardsFromPack()

	if Amount(len(allCards)) < amount {
		return
	}

	for i := Amount(0); i < amount; i++ {
		cardIndex := rand.Intn(len(allCards))
		card := allCards[cardIndex]
		game.cardsInPack[card]--
		game.cardsOnTable[card]++
		allCards = append(allCards[:cardIndex], allCards[cardIndex+1:]...)
	}
}

func (game *game) take5RandomCards(player *player) {
	allCards := game.getAllCardsFromPack()

	for range 5 {
		cardIndex := rand.Intn(len(allCards))
		card := allCards[cardIndex]
		game.cardsInPack[card]--

		switch card {
		case GoodCamel:
			player.herdSize++
		default:
			player.cards[card]++
		}

		allCards = append(allCards[:cardIndex], allCards[cardIndex+1:]...)
	}

}

func (game *game) getAllCardsFromPack() []GoodType {
	allCards := []GoodType{}
	for key, value := range game.cardsInPack {
		for range value {
			allCards = append(allCards, key)
		}
	}
	return allCards
}

func (game *game) changeCurrentPlayer() {
	switch game.currentPlayer.name {
	case game.player1.name:
		game.currentPlayer = game.player2
	case game.player2.name:
		game.currentPlayer = game.player1
	}
}

func getBonus(goodsNumber Amount) Score {
	if goodsNumber < 3 {
		return Score(0)
	}

	if goodsNumber > 5 {
		goodsNumber = 5
	}

	minMax := bonuses[goodsNumber]
	min := int(minMax[0])
	max := int(minMax[1])
	return Score(rand.Intn(max-min)) + minMax[0]
}
