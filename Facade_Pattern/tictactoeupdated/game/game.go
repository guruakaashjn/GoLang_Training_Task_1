package game

import (
	"fmt"
	"tictactoeupdated/board"
	"tictactoeupdated/player"
)

type Game struct {
	Players     [2]*player.Player
	Board       board.Board
	turn        uint
	isGameEnded bool
}

func NewGame(player0Name, player1Name string) *Game {
	var players = [2]*player.Player{player.NewPlayer(player0Name, "X"), player.NewPlayer(player1Name, "O")}
	// fmt.Println("Game start")
	return &Game{

		Players:     players,
		Board:       *board.NewBoard(),
		turn:        0,
		isGameEnded: false,
	}
}

func (g *Game) Play() {
	for i := 0; i < 1; {
	startturn:

		g.Board.PrintBoard()
		fmt.Println("Enter cell number: ")
		var cellNumber uint
		fmt.Scan(&cellNumber)
		var flag, response = g.PlayLogic(cellNumber)
		if flag {
			g.Board.PrintBoard()
			fmt.Println(response)
			i++
			continue
		}
		if !flag {
			if response == "Cell Not Empty" {
				fmt.Println(response)
				goto startturn
			} else {
				g.turn++
				fmt.Println(response)
			}

		}

	}
}

func (g *Game) PlayLogic(cellNumber uint) (bool, string) {

	if !g.Board.IsEmpty(cellNumber) {
		return false, "Cell Not Empty"
	}
	var currentPlayer *player.Player = g.Players[g.turn%2]
	g.Board.MarkCell(cellNumber, currentPlayer.GetSymbol())
	if g.Board.CheckWin() {
		g.isGameEnded = true
		return true, "Player " + currentPlayer.GetName() + " wins with symbol " + currentPlayer.GetSymbol() + "!!!!"
	}

	return false, "Next Turn"

}
