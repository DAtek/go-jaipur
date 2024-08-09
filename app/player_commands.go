package app

import "jaipur/fsm"

type playerCommandCollection map[string]func() fsm.StateName
