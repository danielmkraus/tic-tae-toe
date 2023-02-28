package test

import (
	"github.com/danielmkraus/tictaetoe/game"
	"testing"
)

func TestNewGameCreatedEmpty(t *testing.T) {
	g := game.NewGame()
	if g.Status != game.PLAYER_1_TURN {
		t.Error("Expected to start game with player 1 turn, but is not")
	}
}

func TestPlayer1Winning(t *testing.T) {
	g := game.NewGame()
	playWithoutError(t, g, 0, 0)
	playWithoutError(t, g, 0, 1)
	playWithoutError(t, g, 1, 1)
	playWithoutError(t, g, 0, 2)
	playWithoutError(t, g, 2, 2)

	if g.Status != game.PLAYER_1_WON {
		t.Error("Expecting player 1 to be the winner, but it isn`t")
	}
}

func TestPlayer2Winning(t *testing.T) {
	g := game.NewGame()
	playWithoutError(t, g, 2, 1)
	playWithoutError(t, g, 0, 0)
	playWithoutError(t, g, 0, 1)
	playWithoutError(t, g, 1, 1)
	playWithoutError(t, g, 0, 2)
	playWithoutError(t, g, 2, 2)

	if g.Status != game.PLAYER_2_WON {
		t.Error("Expecting player 2 to be the winner, but it isn`t")
	}
}

func TestWinningInRow(t *testing.T) {
	g := game.NewGame()
	playWithoutError(t, g, 0, 0)
	playWithoutError(t, g, 1, 0)
	playWithoutError(t, g, 0, 1)
	playWithoutError(t, g, 1, 1)
	playWithoutError(t, g, 0, 2)

	if g.Status != game.PLAYER_1_WON {
		t.Error("Expecting player 1 to be the winner, but it isn`t")
	}
}

func TestWinningInColumn(t *testing.T) {
	g := game.NewGame()
	playWithoutError(t, g, 0, 0)
	playWithoutError(t, g, 0, 1)
	playWithoutError(t, g, 1, 0)
	playWithoutError(t, g, 1, 1)
	playWithoutError(t, g, 2, 0)

	if g.Status != game.PLAYER_1_WON {
		t.Error("Expecting player 1 to be the winner, but it isn`t")
	}
}

// X X O
// O O X
// X O X
func TestDraw(t *testing.T) {
	g := game.NewGame()
	playWithoutError(t, g, 0, 0)
	playWithoutError(t, g, 0, 2)
	playWithoutError(t, g, 0, 1)
	playWithoutError(t, g, 1, 0)
	playWithoutError(t, g, 1, 2)
	playWithoutError(t, g, 1, 1)
	playWithoutError(t, g, 2, 0)
	playWithoutError(t, g, 2, 1)
	playWithoutError(t, g, 2, 2)

	if g.Status != game.DRAW {
		t.Error("Expecting player 1 to be the winner, but it isn`t")
	}
}

func playWithoutError(t *testing.T, g *game.Game, c int, r int) {
	err := g.Play(c, r)
	if err != nil {
		t.Errorf("Error while playing %s", err)
	}
}
