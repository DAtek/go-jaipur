package core

const SealsOfExcellenceToWin = Score(2)

func (game *Game) GameEnded() bool {
	return (*game.gameEnded)()
}

func gameEnded(game *Game) bool {
	for _, p := range []*player{game.player1, game.player2} {
		if p.sealsOfExcellence >= SealsOfExcellenceToWin {
			return true
		}
	}

	return false
}
