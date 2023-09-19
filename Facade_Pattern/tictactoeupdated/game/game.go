package game

import (
	"fmt"
	"tictactoeupdated/board"
	"tictactoeupdated/errors"
	"tictactoeupdated/player"
)

type Game struct {
	Players     [2]*player.Player
	board       board.Board
	turn        uint
	isGameEnded bool
}

func NewGame(player0Name, player1Name string) *Game {
	var players = [2]*player.Player{player.NewPlayer(player0Name, "X"), player.NewPlayer(player1Name, "O")}
	// fmt.Println("Game start")
	return &Game{

		Players:     players,
		board:       *board.NewBoard(),
		turn:        0,
		isGameEnded: false,
	}
}

func (g *Game) Play() {
	for i := 0; i < 1; {
	startturn:

		g.board.PrintBoard()

		fmt.Printf("Turn %d\n", g.turn)
		fmt.Println("Enter cell number: ")
		var cellNumber uint
		fmt.Scan(&cellNumber)
		var flag, response = g.PlayLogic(cellNumber)
		if flag {
			g.board.PrintBoard()
			fmt.Println(response)
			g.isGameEnded = true
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

func (g *Game) PlayLogic(cellNumber uint) (flag bool, response string) {
	// var response string
	defer func() {
		if a := recover(); a != nil {

			flag = false
			response = a.(string)
			// fmt.Println("Recovered", response)

		}
	}()

	flag, response = g.PlayLogicInner(cellNumber)
	if flag {
		// fmt.Println(response)
		flag = true
	}
	return flag, response

}

func (g *Game) PlayLogicInner(cellNumber uint) (bool, string) {
	if !g.board.IsEmpty(cellNumber) {
		panic(errors.NewInvalidMove("Cell Not Empty").GetSpecificMessage())
	}
	var currentPlayer *player.Player = g.Players[g.turn%2]
	g.board.MarkCell(cellNumber, currentPlayer.GetSymbol())
	if g.board.CheckWin() {
		g.isGameEnded = true
		return true, "Player " + currentPlayer.GetName() + " wins with symbol " + currentPlayer.GetSymbol() + "!!!!"
	}
	panic(string(errors.NewNextTurn().GetNewNextTurnError()))
}
