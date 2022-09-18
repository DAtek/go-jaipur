package core

func newGame() *Game {
	player1Name := Name("Max")
	player2Name := Name("Martha")
	player1 := player{player1Name, Score(0), GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0, Score(0)}
	player2 := player{player2Name, Score(0), GoodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0, Score(0)}

	return &Game{
		player1:         &player1,
		player2:         &player2,
		cardsOnTable:    GoodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
		cardsInPack:     GoodMap{GoodCloth: Amount(50)},
		currentPlayer:   &player1,
		soldGoods:       GoodMap{},
		resetAfterRound: func() {},
		roundEnded:      func() bool { return false },
		gameEnded:       func() bool { return false },
	}
}
