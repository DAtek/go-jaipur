package core

func (game *game) RoundWinner() (Name, error) {
	return game.roundWinner()
}

func roundWinner(game *game) (Name, error) {
	if !game.roundEnded() {
		return "", RoundNotEndedError
	}

	if game.player1.herdSize > game.player2.herdSize {
		game.player1.score += 5
	}

	if game.player2.herdSize > game.player1.herdSize {
		game.player2.score += 5
	}

	if game.player1.score > game.player2.score {
		return game.player1.name, nil
	}

	if game.player2.score > game.player1.score {
		return game.player2.name, nil
	}
	return Name(EmptyName), nil
}
