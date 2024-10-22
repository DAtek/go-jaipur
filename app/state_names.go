package app

import "github.com/DAtek/fsm"

const (
	STATE_GAME_STARTING = fsm.StateName("Ask for names")
	STATE_PLAYER_TURN   = fsm.StateName("Player's turn")
	STATE_ROUND_ENDED   = fsm.StateName("Round ended")
	STATE_GAME_ENDED    = fsm.StateName("Game ended")
)
