package fsm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFSM(t *testing.T) {
	t.Run("FSM returns final state's name", func(t *testing.T) {
		transit1 := func() StateName {
			return "state2"
		}
		state1 := &State{
			Name:    "state1",
			Variant: VariantStart,
			Transit: &transit1,
		}

		transit2 := func() StateName {
			return "state3"
		}
		state2 := &State{
			Name:    "state2",
			Variant: VariantIntermediate,
			Transit: &transit2,
		}

		state3 := &State{
			Name:    "state3",
			Variant: VariantFinal,
		}

		fsm := FSM{
			States: []*State{state1, state2, state3},
		}

		result := fsm.Run()

		assert.Equal(t, result, state3.Name)
	})

	t.Run("FSM executes transitions", func(t *testing.T) {
		type context struct {
			count uint8
		}

		c := context{0}
		transit1 := func() StateName {
			c.count++
			return "state2"
		}
		state1 := &State{
			Name:    "state1",
			Variant: VariantStart,
			Transit: &transit1,
		}

		transit2 := func() StateName {
			c.count += 2
			return "state3"
		}
		state2 := &State{
			Name:    "state2",
			Variant: VariantIntermediate,
			Transit: &transit2,
		}

		state3 := &State{
			Name:    "state3",
			Variant: VariantFinal,
		}

		fsm := FSM{
			States: []*State{state1, state2, state3},
		}

		fsm.Run()

		assert.Equal(t, uint8(3), c.count)
	})
}
