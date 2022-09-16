package core

const SameNamesError = JaipurError("Players have the same name")

func NewGame(player1Name, player2Name Name) (*game, error) {
	if player1Name == player2Name {
		return nil, error(SameNamesError)
	}

	player1 := &player{player1Name, Score(0), goodMap{}, 0}
	player2 := &player{player2Name, Score(0), goodMap{}, 0}

	players := map[Name]*player{
		player1Name: player1,
		player2Name: player2,
	}

	game := game{
		players:      players,
		soldGoods:    goodMap{},
		cardsOnTable: goodMap{},
		cardsInPack:  goodMap{},
	}

	for key, value := range allCards {
		game.cardsInPack[key] = value
	}

	game.cardsOnTable[GoodCamel] = 3
	game.moveCardsFromPackToTable(2)

	return &game, nil
}
