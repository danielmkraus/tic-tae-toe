package game

type TicTaeToeRepository map[string]*Game

func (t TicTaeToeRepository) insert(game *Game) {
	t[game.id.String()] = game
}

func (t TicTaeToeRepository) delete(game *Game) {
	t[game.id.String()] = nil
}
