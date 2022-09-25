package app

import (
	"fmt"
	"jaipur/fsm"
)

func finishRound(app *App) fsm.StateName {
	winner, _ := app.game.RoundWinner()
	fmt.Fprint(app.writer, "Winner of the round: "+string(winner)+"\n")
	playerScores := app.game.PlayerScores()
	for name, score := range playerScores {
		fmt.Fprintf(app.writer, "%s's score: %d\n", name, score)
	}
	input(app.reader, app.writer, "Press enter to continue")
	err := app.game.FinishRound()

	if err != nil {
		panic("Unexpected error: " + err.Error())
	}

	fmt.Fprint(app.writer, clearScreenString)

	if app.game.GameEnded() {
		return gameEnded.Name
	}

	return playerTurn.Name
}
