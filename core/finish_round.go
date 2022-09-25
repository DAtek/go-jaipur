package core

const RoundNotEndedError = JaipurError("Round not ended")

func (game *Game) FinishRound() error {
	if game.gameEnded() {
		return GameEndedError
	}

	if !game.roundEnded() {
		return RoundNotEndedError
	}

	winner, _ := (game.roundWinner)()

	for _, p := range []*player{game.player1, game.player2} {
		if p.name == winner {
			p.sealsOfExcellence++
			break
		}
	}

	game.resetAfterRound()

	return nil
}
