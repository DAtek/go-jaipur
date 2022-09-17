package core

const SameNamesError = JaipurError("Players have the same name")

func NewGame(player1Name, player2Name Name) (*game, error) {
	if player1Name == player2Name {
		return nil, error(SameNamesError)
	}

	player1 := &player{player1Name, Score(0), goodMap{}, 0, Score(0)}

	game := game{
		player1:       player1,
		player2:       &player{player2Name, Score(0), goodMap{}, 0, Score(0)},
		soldGoods:     goodMap{},
		cardsOnTable:  goodMap{},
		cardsInPack:   goodMap{},
		currentPlayer: player1,
	}

	for key, value := range cardsInGame {
		game.cardsInPack[key] = value
	}

	game.cardsOnTable[GoodCamel] = 3
	game.cardsInPack[GoodCamel] -= 3

	game.moveCardsFromPackToTable(2)
	game.take5RandomCards(player1)
	game.take5RandomCards(game.player2)

	return &game, nil
}
