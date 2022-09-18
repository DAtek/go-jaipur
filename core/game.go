package core

type IGame interface {
	CurrentPlayerName() Name
	CurrentPlayerCards() GoodMap
	CardsOnTable() GoodMap
	TakeCard(card GoodType) error
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
	cards := GoodMap{}
	for k, v := range game.currentPlayer.cards {
		cards[k] = v
	}
	return cards
}

func (game *Game) CardsOnTable() GoodMap {
	cards := GoodMap{}
	for k, v := range game.cardsOnTable {
		cards[k] = v
	}
	return cards
}

func (game *Game) CurrentPlayerName() Name {
	return game.currentPlayer.name
}
