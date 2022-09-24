package app

import "jaipur/fsm"

type playerCommandCollection struct {
	Buy      *func() fsm.StateName
	Exchange *func() fsm.StateName
	Sell     *func() fsm.StateName
}
