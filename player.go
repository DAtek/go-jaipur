package jaipur

type player struct {
	name  Name
	score Score
	cards productMap
}

type playerMap map[Name]*player
