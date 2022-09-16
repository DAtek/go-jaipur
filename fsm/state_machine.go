package fsm

type FSM struct {
	states []State
}

func (fsm *FSM) Run() StateName {
	var currentState *State
	for _, state := range fsm.states {
		if state.variant == VariantStart {
			currentState = &state
			break
		}
	}

	for currentState.variant != VariantFinal {
		nextStateName := currentState.transit()
		for _, state := range fsm.states {
			if state.name == nextStateName {
				currentState = &state
				break
			}
		}
	}

	return currentState.name
}
