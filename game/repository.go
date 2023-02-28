package game

type TicTaeToeRepository map[string]Game

func (t TicTaeToeRepository) insert(game *Game) {
	t[game.Id.String()] = *game
}

func (t TicTaeToeRepository) delete(game Game) {
	delete(t, game.Id.String())
}

func (t TicTaeToeRepository) get(gameId string) Game {
	return t[gameId]
}
