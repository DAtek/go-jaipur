package fsm

type FSM struct {
	States []*State
}

func (fsm *FSM) Run() StateName {
	var currentState *State
	for _, state := range fsm.States {
		if state.Variant == VariantStart {
			currentState = state
			break
		}
	}

	for currentState.Variant != VariantFinal {
		nextStateName := currentState.Transit()
		for _, state := range fsm.States {
			if state.Name == nextStateName {
				currentState = state
				break
			}
		}
	}

	return currentState.Name
}
