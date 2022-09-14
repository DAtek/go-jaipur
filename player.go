package jaipur

type player struct {
	name     Name
	score    Score
	cards    goodMap
	herdSize Amount
}

type playerMap map[Name]*player
