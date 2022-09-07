package jaipur

import (
	"math/rand"
	"time"
)

func (game *game) moveCardsFromPackToTable(amount Amount) {
	allCards_ := []ProductType{}
	for key, value := range game.cardsInPack {
		for i := Amount(0); i < value; i++ {
			allCards_ = append(allCards_, key)
		}
	}

	for i := Amount(0); i < amount; i++ {
		rand.Seed(time.Now().UnixNano())
		cardIndex := rand.Intn(len(allCards_))
		card := allCards_[cardIndex]
		game.cardsInPack[card]--
		game.cardsOnTable[card]++
		allCards_ = append(allCards_[:cardIndex], allCards_[cardIndex+1:]...)
	}
}

func getBonus(goodsNumber Amount) Score {
	if goodsNumber < 3 {
		return Score(0)
	}

	if goodsNumber > 5 {
		goodsNumber = 5
	}

	rand.Seed(time.Now().UnixNano())
	minMax := bonuses[goodsNumber]
	min := int(minMax[0])
	max := int(minMax[1])
	return Score(rand.Intn(max-min)) + minMax[0]
}
