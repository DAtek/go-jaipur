package core

type IGame interface {
	CurrentPlayerName() Name
	CurrentPlayerCards() GoodMap
	CardsOnTable() GoodMap
	Buy(card GoodType) error
	Sell(card GoodType) error
	Exchange(buy, sell GoodMap) error
	RoundEnded() bool
	GameEnded() bool
	RoundWinner() (Name, error)
	PlayerScores() ScoreMap
	FinishRound() error
	GameWinner() (Name, error)
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
	roundWinner     func() (Name, error)
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

func (game *Game) PlayerScores() ScoreMap {
	return ScoreMap{
		game.player1.name: game.player1.score,
		game.player2.name: game.player2.score,
	}
}

const GameNotEndedError = JaipurError("Game not ended")

func (game *Game) GameWinner() (Name, error) {
	if !game.gameEnded() {
		return "", GameNotEndedError
	}
	if game.player1.sealsOfExcellence > game.player2.sealsOfExcellence {
		return game.player1.name, nil
	} else {
		return game.player2.name, nil
	}
}
