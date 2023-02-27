package game

import (
	"github.com/google/uuid"
	"time"
)

type Game struct {
	id          uuid.UUID
	start       time.Time
	end         time.Time
	board       [][]int
	computerWon bool
}

func NewGame() *Game {
	return &Game{
		id:    uuid.New(),
		start: time.Now(),
		board: [][]int{{0, 0, 0}, {0, 0, 0}, {0, 0, 0}},
	}
}

func (g *Game) Play(column int, row int) bool {
	g.board[column][row] = 1
	return false
}

func (g *Game) isFinished() bool {

	return g.board[0][0] != 0 &&
		(g.board[0][0] == g.board[0][1] && g.board[0][0] == g.board[0][2]) ||
		(g.board[0][0] == g.board[1][1] && g.board[0][0] == g.board[2][2])

}
