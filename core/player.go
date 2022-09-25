package core

const MaximumCardsInHand = Amount(7)

type player struct {
	name              Name
	score             Score
	cards             GoodMap
	herdSize          Amount
	sealsOfExcellence Score
}

func (player *player) resetAfterRound() {
	player.score = 0
	player.cards = GoodMap{}
	player.herdSize = 0
}
