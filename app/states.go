package app

import "jaipur/fsm"

var gameStart = fsm.State{
	Name:    "Ask for names",
	Variant: fsm.VariantStart,
}

var playerTurn = fsm.State{
	Name:    "Player's turn",
	Variant: fsm.VariantIntermediate,
}

var roundEnded = fsm.State{
	Name:    "Round ended",
	Variant: fsm.VariantIntermediate,
}

var gameEnded = fsm.State{
	Name:    "Game ended",
	Variant: fsm.VariantIntermediate,
}

var finalState = fsm.State{
	Name:    "The end",
	Variant: fsm.VariantFinal,
}
