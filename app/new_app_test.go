package app

import (
	"jaipur/fsm"
	"strings"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewApp(t *testing.T) {
	t.Run("Simulate a whole play", func(t *testing.T) {
		sim := simulator{
			player1: "Carla",
			player2: "Jenna",
			reader:  make(chan string, 1),
			writer:  make(chan string, 1),
			t:       t,
		}

		app := NewApp(&sim.reader, &sim.writer)
		wg := sync.WaitGroup{}
		wg.Add(2)

		go func() {
			app.Run()
			wg.Done()
		}()

		go func() {
			sim.run()
			wg.Done()
		}()

		wg.Wait()

		assert.Equal(t, sim.player1, *sim.winner)
	})
}

type channelReaderWriter chan string

func (ch *channelReaderWriter) Read(p []byte) (int, error) {
	s := <-*ch
	in := []byte(s)
	in = append(in, '\n')
	copy(p, in)
	return len(p), nil
}

func (ch *channelReaderWriter) Write(p []byte) (int, error) {
	s := string(p)
	*ch <- s
	return len(p), nil
}

type simulator struct {
	player1 string
	player2 string
	winner  *string
	reader  channelReaderWriter
	writer  channelReaderWriter
	t       *testing.T
}

func (s *simulator) run() {
	s.enterPlayerNames()

	player1Turn := fsm.State{
		Name:    "player 1",
		Variant: fsm.VariantStart,
		Transit: func() fsm.StateName { return s.alwaysBuyOrSell() },
	}

	player2Turn := fsm.State{
		Name:    "player 2",
		Variant: fsm.VariantIntermediate,
		Transit: func() fsm.StateName { return s.alwaysExchange() },
	}

	showWinner := fsm.State{
		Name:    "show winner",
		Variant: fsm.VariantIntermediate,
		Transit: func() fsm.StateName { return s.showWinner() },
	}

	finalState := fsm.State{
		Name:    "final",
		Variant: fsm.VariantFinal,
	}

	stateMachine := fsm.FSM{
		States: []*fsm.State{&player1Turn, &player2Turn, &showWinner, &finalState},
	}

	stateMachine.Run()
}

func (s *simulator) enterPlayerNames() {
	s.waitForOutput("player 1")
	s.simulateInput(s.player1)
	s.waitForOutput("player 2")
	s.simulateInput(s.player2)
}

func (s *simulator) alwaysBuyOrSell() fsm.StateName {
	s.t.Log("PLAYER 1 STATE **********\n")
	for _, action := range []string{"B", "S"} {
		for good := range goodAbbreviations {
			s.waitForOutput("Pick an action")
			s.simulateInput(action)
			switch action {
			case "B":
				s.waitForOutput("buy:")
			case "S":
				s.waitForOutput("sell:")
			}
			s.simulateInput(good)
			out := s.readOutput()
			if strings.Contains(out, s.player2) {
				return "player 2"
			}
			if strings.Contains(out, "Winner") {
				return "show winner"
			}
		}
	}

	panic("Player 1 couldn't complete the turn")
}

func (s *simulator) alwaysExchange() fsm.StateName {
	s.t.Log("PLAYER 2 STATE **********\n")
	for good1 := range goodAbbreviations {
		for good2 := range goodAbbreviations {
			s.waitForOutput("Pick an action")
			s.simulateInput("E")
			s.waitForOutput("Buy")
			s.simulateInput("1" + good1)

			if output := s.readOutput(); !strings.Contains(output, "Sell") {
				continue
			}

			s.simulateInput("1" + good2)
			if output := s.readOutput(); strings.Contains(output, s.player1) {
				return "player 1"
			}
		}
	}

	panic("Player 2 couldn't successfuly exchange")
}

func (s *simulator) showWinner() fsm.StateName {
	s.t.Log("SHOW WINNER STATE **********\n")
	s.waitForOutput("Press")
	s.simulateInput("enter")
	output := s.readOutput()
	if strings.Contains(output, "Congratulations") {
		s.simulateInput("enter")
		if strings.Contains(output, s.player1) {
			s.winner = &s.player1
		} else {
			s.winner = &s.player2
		}
		return "final"
	}
	if strings.Contains(output, s.player1) {
		return "player 1"
	}
	if strings.Contains(output, s.player2) {
		return "player 2"
	}
	panic("WRONG HANDLING OF FINAL STATE")
}

func (s *simulator) waitForOutput(wanted string) string {
	for {
		if out := s.readOutput(); strings.Contains(out, wanted) {
			return out
		}
	}
}

func (s *simulator) readOutput() string {
	out := readFirstNonemptyLine(s.writer)
	s.t.Logf("OUTPUT: %s", out)
	return out
}

func (s *simulator) simulateInput(input string) {
	s.t.Logf("INPUT : %s", input)
	s.reader <- input
}

func readFirstNonemptyLine(c chan string) string {
	for {
		out := <-c
		switch out {
		case clearScreenString:
		case "\n":
		case "":
		case " ":
		default:
			return out
		}
	}
}
