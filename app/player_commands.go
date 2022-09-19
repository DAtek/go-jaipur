package app

import "jaipur/fsm"

type playerCommandCollection struct {
	TakeCard      func() fsm.StateName
	ExchangeCards func() fsm.StateName
	SellCards     func() fsm.StateName
}
