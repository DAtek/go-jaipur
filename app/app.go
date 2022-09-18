package app

import (
	"io"
	"jaipur/core"
	"jaipur/fsm"
)

type App struct {
	game   core.IGame
	fsm    *fsm.FSM
	reader io.Reader
	writer io.Writer
}
