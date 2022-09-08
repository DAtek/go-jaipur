package jaipur

type player struct {
	name     Name
	score    Score
	cards    productMap
	herdSize Amount
}

type playerMap map[Name]*player
