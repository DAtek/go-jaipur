package core

type IGame interface {
	CurrentPlayerName() Name
	CurrentPlayerCards() GoodMap
	CardsOnTable() GoodMap
	TakeCard(card GoodType) error
	SellGoods(card GoodType) error
	RoundEnded() bool
}

type Game struct {
	player1         *player
	player2         *player
	soldGoods       GoodMap
	cardsInPack     GoodMap
	cardsOnTable    GoodMap
	currentPlayer   *player
	resetAfterRound func()
	roundEnded      func() bool
	gameEnded       func() bool
}

func (game *Game) CurrentPlayerCards() GoodMap {
	cards := game.currentPlayer.cards.Copy()
	cards[GoodCamel] = game.currentPlayer.herdSize
	return cards
}

func (game *Game) CardsOnTable() GoodMap {
	return game.cardsOnTable.Copy()
}

func (game *Game) CurrentPlayerName() Name {
	return game.currentPlayer.name
}
