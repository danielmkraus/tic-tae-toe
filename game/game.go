package game

import (
	"errors"
	"github.com/google/uuid"
	"time"
)

type Game struct {
	Id     uuid.UUID
	Start  time.Time
	End    time.Time
	Board  [][]GameStatus
	Status GameStatus
}

type GameStatus int

const (
	PLAYER_1_TURN GameStatus = -1
	PLAYER_2_TURN GameStatus = -2
	DRAW          GameStatus = 0
	PLAYER_1_WON  GameStatus = 1
	PLAYER_2_WON  GameStatus = 2
)

func NewGame() *Game {
	return &Game{
		Id:     uuid.New(),
		Start:  time.Now(),
		Board:  [][]GameStatus{{DRAW, DRAW, DRAW}, {DRAW, DRAW, DRAW}, {DRAW, DRAW, DRAW}},
		Status: PLAYER_1_TURN,
	}
}

func (g *Game) Play(column int, row int) error {
	if g.Status >= DRAW {
		return errors.New("game is finished")
	}
	if g.Board[column][row] != DRAW {
		return errors.New("position already played")
	}
	g.Board[column][row] = g.Status * -1
	g.Status = g.checkGameStatus()
	if g.Status >= DRAW {
		g.End = time.Now()
	}
	return nil
}

func (g *Game) checkGameStatus() GameStatus {
	// check diagonals
	c := g.Board[1][1]
	if c != DRAW && ((c == g.Board[0][0] && c == g.Board[2][2]) ||
		(c == g.Board[0][2] && c == g.Board[2][0])) {
		return c
	}

	playedPositions := 0
	for i, r := range g.Board {
		// check if row is played and equal values are set
		c := r[0]
		if c != DRAW && c == r[1] && c == r[2] {
			return c
		}
		// check if column is played and equal values are set
		if c != DRAW && g.Board[1][i] == c && g.Board[2][i] == c {
			return c
		}
		// count how many positions has been played
		for _, v := range r {
			if v != DRAW {
				playedPositions++
			}
		}
	}
	if playedPositions == 9 {
		return DRAW
	}
	if playedPositions%2 == 0 {
		return PLAYER_1_TURN
	}
	return PLAYER_2_TURN
}
