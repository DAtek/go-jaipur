package core

const SameNamesError = JaipurError("Players have the same name")
const EmptyNameError = JaipurError("Player got empty name")

const EmptyName = ""

func NewGame(player1Name, player2Name Name) (IGame, error) {
	if player1Name == EmptyName || player2Name == EmptyName {
		return nil, EmptyNameError
	}

	if player1Name == player2Name {
		return nil, error(SameNamesError)
	}

	game := &game{
		player1:     &player{name: player1Name, sealsOfExcellence: 0},
		player2:     &player{name: player2Name, sealsOfExcellence: 0},
		cardsInPack: GoodMap{},
	}

	game.roundEnded = func() bool { return roundEnded(game) }
	game.gameEnded = func() bool { return gameEnded(game) }
	game.roundWinner = func() (Name, error) { return roundWinner(game) }
	game.currentPlayer = game.player1
	game.resetAfterRound = func() {
		for key, value := range cardsInGame {
			game.cardsInPack[key] = value
		}
		game.player1.resetAfterRound()
		game.player2.resetAfterRound()
		game.cardsOnTable = GoodMap{}
		game.cardsOnTable[GoodCamel] = 3
		game.cardsInPack[GoodCamel] -= 3
		game.soldGoods = GoodMap{}

		game.moveCardsFromPackToTable(2)
		game.take5RandomCards(game.player1)
		game.take5RandomCards(game.player2)
	}
	game.resetAfterRound()

	return game, nil
}
