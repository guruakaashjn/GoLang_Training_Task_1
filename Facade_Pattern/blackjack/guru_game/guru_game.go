package guru_game

import (
	"blackjack/errors"
	"blackjack/guru_player"
	"fmt"
)

type Game struct {
	players    [2]*guru_player.Player
	turn       int
	isGameOver bool
}

func NewGame(player1Name, player2Name string) *Game {
	var players = [2]*guru_player.Player{guru_player.NewPlayer(player1Name), guru_player.NewPlayer(player2Name)}
	return &Game{
		players:    players,
		turn:       0,
		isGameOver: false,
	}
}

func (g *Game) Play() {
	var singleFlag int = 0
	for i := 0; i < 1; {
		var currentPlayer *guru_player.Player = g.players[g.turn%2]

		g.PrintCards()

		// if g.players[0].GetHold() && g.players[1].GetHold(){

		// }

		single := g.playInner()
		if !single {
			fmt.Printf("Turn of Player %s\n", g.players[g.turn%2].GetPlayerName())
		}

		if !g.isGameOver {

			fmt.Println("1. Draw a card")
			fmt.Println("2. Hold cards")
			var choice int
			fmt.Scan(&choice)

			switch choice {
			case 1:
				currentPlayer.AddCardToDeck()

			case 2:
				currentPlayer.SetHold()
				singleFlag = 1

			}
		} else {
			g.checkWin()
			i++
		}

		if singleFlag < 2 {
			g.turn++
		}
		if singleFlag == 1 {
			singleFlag = 2
		}

	}

}

func (g *Game) checkWin() {
	cardsSumPlayer1 := g.players[0].SumOfCards()
	cardsSumPlayer2 := g.players[1].SumOfCards()
	fmt.Printf("\nSum of cards of player 1: %d", cardsSumPlayer1)
	fmt.Printf("\nSum of cards of player 2: %d", cardsSumPlayer2)
	if cardsSumPlayer1 > 21 && cardsSumPlayer2 > 21 {
		fmt.Printf("\nNo player wins the game.")
	}
	if cardsSumPlayer1 > 21 {
		fmt.Printf("\nPlayer %s wins with %d points", g.players[1].GetPlayerName(), cardsSumPlayer2)
	}
	if cardsSumPlayer2 > 21 {
		fmt.Printf("\nPlayer %s wins with %d points", g.players[0].GetPlayerName(), cardsSumPlayer1)
	}
	if cardsSumPlayer1 > cardsSumPlayer2 && cardsSumPlayer1 <= 21 {
		fmt.Printf("\nPlayer %s wins with %d points", g.players[0].GetPlayerName(), cardsSumPlayer1)
	}
	if cardsSumPlayer1 < cardsSumPlayer2 && cardsSumPlayer2 <= 21 {
		fmt.Printf("\nPlayer %s wins with %d points", g.players[1].GetPlayerName(), cardsSumPlayer2)
	}
}

func (g *Game) playInner() (single bool) {

	defer func() {
		if a := recover(); a != nil {
			fmt.Println(a)
		}
	}()

	if g.players[0].GetHold() && g.players[1].GetHold() {
		g.isGameOver = true
		single = true

		panic(errors.NewDrawError("\nBoth Players Made Hold").GetSpecificMessage())
	}
	if g.players[0].GetHold() || g.players[1].GetHold() {
		msg := "\nPlayer " + g.players[(g.turn+1)%2].GetPlayerName() + " Made Hold"
		panic(errors.NewDrawError(msg).GetSpecificMessage())
	} else {
		panic(errors.NewDrawError("\nNo Player Made Hold").GetSpecificMessage())
	}

}

func (g *Game) PrintCards() {
	for i := 0; i < 2; i++ {
		g.players[i].PrintCards()
		fmt.Println()
	}
}
