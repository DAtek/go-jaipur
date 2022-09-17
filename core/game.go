package core

type game struct {
	player1         *player
	player2         *player
	soldGoods       goodMap
	cardsInPack     goodMap
	cardsOnTable    goodMap
	currentPlayer   *player
	resetAfterRound func()
	roundEnded      func() bool
	gameEnded       func() bool
}
