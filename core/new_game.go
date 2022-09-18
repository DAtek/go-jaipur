package core

const SameNamesError = JaipurError("Players have the same name")

func NewGame(player1Name, player2Name Name) (*Game, error) {
	if player1Name == player2Name {
		return nil, error(SameNamesError)
	}

	game := &Game{
		player1:      &player{name: player1Name, sealsOfExcellence: 0},
		player2:      &player{name: player2Name, sealsOfExcellence: 0},
		soldGoods:    GoodMap{},
		cardsOnTable: GoodMap{},
		cardsInPack:  GoodMap{},
	}

	game.roundEnded = func() bool {
		return roundEnded(game)
	}

	game.gameEnded = func() bool {
		return gameEnded(game)
	}

	game.currentPlayer = game.player1

	resetAfterRound := func() {
		for key, value := range cardsInGame {
			game.cardsInPack[key] = value
		}
		game.player1.resetAfterRound()
		game.player2.resetAfterRound()
		game.cardsOnTable[GoodCamel] = 3
		game.cardsInPack[GoodCamel] -= 3

		game.moveCardsFromPackToTable(2)
		game.take5RandomCards(game.player1)
		game.take5RandomCards(game.player2)
	}

	game.resetAfterRound = resetAfterRound
	game.resetAfterRound()

	return game, nil
}
