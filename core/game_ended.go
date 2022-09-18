package core

func (game *Game) GameEnded() bool {
	return game.gameEnded()
}

func gameEnded(game *Game) bool {
	return game.player1.sealsOfExcellence == 2 || game.player2.sealsOfExcellence == 2
}
