package app

import (
	"io"
	"jaipur/core"

	"github.com/DAtek/fsm"
)

type App struct {
	Game           core.IGame
	Reader         io.Reader
	Writer         io.Writer
	PlayerCommands PlayerCommandCollection
}

type PlayerCommandCollection map[string]func() fsm.StateName
