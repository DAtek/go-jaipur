package core

func newGame() *Game {
	player1Name := Name("Max")
	player2Name := Name("Martha")
	player1 := player{player1Name, Score(0), GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0, Score(0)}
	player2 := player{player2Name, Score(0), GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0, Score(0)}

	game := Game{
		player1:       &player1,
		player2:       &player2,
		cardsOnTable:  GoodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
		cardsInPack:   GoodMap{GoodCloth: Amount(50)},
		currentPlayer: &player1,
		soldGoods:     GoodMap{},
	}
	gameEnded := func() bool { return false }
	roundEnded := func() bool { return false }
	resetAfterRound := func() {}
	roundWinner := func() (Name, error) {
		return "", nil
	}
	game.roundWinner = &roundWinner
	game.gameEnded = &gameEnded
	game.roundEnded = &roundEnded
	game.resetAfterRound = &resetAfterRound

	return &game
}
