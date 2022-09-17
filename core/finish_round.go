package core

const RoundNotEndedError = JaipurError("Round not ended")

func (game *game) FinishRound() error {
	if game.gameEnded() {
		return GameEndedError
	}

	if !game.roundEnded() {
		return RoundNotEndedError
	}

	if game.player1.herdSize > game.player2.herdSize {
		game.player1.score += coins[GoodCamel][0]
	}

	if game.player2.herdSize > game.player1.herdSize {
		game.player2.score += coins[GoodCamel][0]
	}

	if game.player1.score > game.player2.score {
		game.player1.sealsOfExcellence++
	}

	if game.player2.score > game.player1.score {
		game.player2.sealsOfExcellence++
	}

	game.resetAfterRound()

	return nil
}
