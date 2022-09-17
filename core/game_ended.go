package core

func (game *game) GameEnded() bool {
	return game.gameEnded()
}

func gameEnded(game *game) bool {
	return game.player1.sealsOfExcellence == 2 || game.player2.sealsOfExcellence == 2
}
