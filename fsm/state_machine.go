package fsm

type IFSM interface {
	Run() StateName
}

type fsm struct {
	states map[StateName]*State
}

func NewFSM(states []*State) IFSM {
	stateMap := map[StateName]*State{}
	for _, state := range states {
		stateMap[state.Name] = state
	}

	return &fsm{states: stateMap}
}

func (f *fsm) Run() StateName {
	var currentState *State
	for _, state := range f.states {
		if state.Variant == VariantStart {
			currentState = state
			break
		}
	}

	for currentState.Variant != VariantFinal {
		currentState = f.states[currentState.Transit()]
	}

	return currentState.Name
}
