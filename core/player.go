package core

type player struct {
	name              Name
	score             Score
	cards             goodMap
	herdSize          Amount
	sealsOfExcellence Score
}

func (player *player) resetAfterRound() {
	player.score = 0
	player.cards = goodMap{}
	player.herdSize = 0
}
