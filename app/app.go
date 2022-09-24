package app

import (
	"bufio"
	"io"
	"jaipur/core"
	"jaipur/fsm"
)

type App struct {
	game   core.IGame
	fsm    *fsm.FSM
	reader *bufio.Reader
	writer io.Writer
}
