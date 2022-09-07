package jaipur

const SameNamesError = JaipurError("Players have the same name")

func NewGame(player1Name, player2Name Name) (*game, error) {
	if player1Name == player2Name {
		return nil, error(SameNamesError)
	}

	player1 := &player{player1Name, Score(0), productMap{}}
	player2 := &player{player2Name, Score(0), productMap{}}

	players := map[Name]*player{
		player1Name: player1,
		player2Name: player2,
	}

	game := game{
		players:      players,
		soldProducts: productMap{},
		cardsOnTable: productMap{},
		cardsInPack:  productMap{},
	}

	for key, value := range allCards {
		game.cardsInPack[key] = value
	}

	game.cardsOnTable[ProductCamel] = 3
	game.moveCardsFromPackToTable(2)

	return &game, nil
}
