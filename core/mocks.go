package core

func newGame() *game {
	player1Name := Name("Max")
	player2Name := Name("Martha")
	player1 := player{player1Name, Score(0), goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0}
	player2 := player{player2Name, Score(0), goodMap{GoodCloth: Amount(2), GoodLeather: Amount(1)}, 0}

	return &game{
		player1:       &player1,
		player2:       &player2,
		cardsOnTable:  goodMap{GoodDiamond: Amount(2), GoodGold: Amount(1)},
		cardsInPack:   goodMap{GoodCloth: Amount(50)},
		currentPlayer: &player1,
		soldGoods:     goodMap{},
	}
}
