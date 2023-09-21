package guru_game

import (
	"blackjack/errors"
	"blackjack/guru_player"
	"fmt"
	"strconv"
)

type Game struct {
	players    [2]*guru_player.Player
	turn       int
	isGameOver bool
	gameResult string
}

func (g *Game) GetTurn() string {
	return "Turn of Player " + g.players[g.turn%2].GetPlayerName() + "\n"
}

func NewGame(player1Name, player2Name string) (*Game, string) {
	var players = [2]*guru_player.Player{guru_player.NewPlayer(player1Name), guru_player.NewPlayer(player2Name)}
	var initialTurn string = "\nTurn of Player " + players[(0)%2].GetPlayerName() + "\n"
	return &Game{
		players:    players,
		turn:       0,
		isGameOver: false,
		gameResult: "",
	}, initialTurn
}

var singleFlag int = 0

func (g *Game) Play(choice int) (flag bool, response string, bothPlayersCardsInHand string) {
	defer func() {
		bothPlayersCardsInHand = g.PrintCards()
	}()

	if !g.isGameOver {

		var currentPlayer *guru_player.Player = g.players[g.turn%2]

		switch choice {
		case 1:
			currentPlayer.AddCardToDeck()

		case 2:
			currentPlayer.SetHold()
			singleFlag = 1
		}

		if singleFlag < 2 {
			g.turn++
		}
		if singleFlag == 1 {
			singleFlag = 2
		}

		single := g.playInner()

		if !single {
			response = "Turn of Player " + g.players[(g.turn)%2].GetPlayerName() + "\n"
			return false, response, bothPlayersCardsInHand
			// fmt.Printf("Turn of Player %s\n", g.players[g.turn%2].GetPlayerName())
		}
		sum1, sum2 := g.checkWin()
		// fmt.Println("sum1:", sum1)
		// fmt.Println("sum2:", sum2)
		response = sum1 + sum2 + g.gameResult
		g.gameResult = sum1 + sum2 + g.gameResult

		return true, response, bothPlayersCardsInHand

	}
	return true, g.gameResult, bothPlayersCardsInHand

}

func (g *Game) playInner() (single bool) {
	single = false
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

func (g *Game) checkWin() (cardsSumPlayer1String, cardsSumPlayer2String string) {
	cardsSumPlayer1 := g.players[0].SumOfCards()
	cardsSumPlayer2 := g.players[1].SumOfCards()
	cardsSumPlayer1String = "\nSum of cards of player 1: " + strconv.Itoa(cardsSumPlayer1)
	cardsSumPlayer2String = "\nSum of cards of player 2: " + strconv.Itoa(cardsSumPlayer2)
	// fmt.Printf("\nSum of cards of player 1: %d", cardsSumPlayer1)
	// fmt.Printf("\nSum of cards of player 2: %d", cardsSumPlayer2)

	if cardsSumPlayer1 > 21 && cardsSumPlayer2 > 21 {
		g.gameResult = "\nNo player wins the game."
		return cardsSumPlayer1String, cardsSumPlayer2String
		// fmt.Printf("\nNo player wins the game.")
	}
	if cardsSumPlayer1 > 21 {
		g.gameResult = "\nPlayer " + g.players[1].GetPlayerName() + " wins with " + strconv.Itoa(cardsSumPlayer2) + " points."
		return cardsSumPlayer1String, cardsSumPlayer2String
		// fmt.Printf("\nPlayer %s wins with %d points", g.players[1].GetPlayerName(), cardsSumPlayer2)
	}
	if cardsSumPlayer2 > 21 {
		g.gameResult = "\nPlayer " + g.players[0].GetPlayerName() + " wins with " + strconv.Itoa(cardsSumPlayer1) + " points."
		return cardsSumPlayer1String, cardsSumPlayer2String
		// fmt.Printf("\nPlayer %s wins with %d points", g.players[0].GetPlayerName(), cardsSumPlayer1)
	}
	if cardsSumPlayer1 > cardsSumPlayer2 && cardsSumPlayer1 <= 21 {
		g.gameResult = "\nPlayer " + g.players[0].GetPlayerName() + " wins with " + strconv.Itoa(cardsSumPlayer1) + " points."
		return cardsSumPlayer1String, cardsSumPlayer2String
		// fmt.Printf("\nPlayer %s wins with %d points", g.players[0].GetPlayerName(), cardsSumPlayer1)
	}
	if cardsSumPlayer1 < cardsSumPlayer2 && cardsSumPlayer2 <= 21 {
		g.gameResult = "\nPlayer " + g.players[1].GetPlayerName() + " wins with " + strconv.Itoa(cardsSumPlayer2) + " points."
		return cardsSumPlayer1String, cardsSumPlayer2String
		// fmt.Printf("\nPlayer %s wins with %d points", g.players[1].GetPlayerName(), cardsSumPlayer2)
	}
	if cardsSumPlayer1 == cardsSumPlayer2 {
		g.gameResult = "\nPlayer " + g.players[0].GetPlayerName() + " and Player " + g.players[1].GetPlayerName() + " wins with " + strconv.Itoa(cardsSumPlayer2) + " points."
		return cardsSumPlayer1String, cardsSumPlayer2String
	}
	return cardsSumPlayer1String, cardsSumPlayer2String

}

func (g *Game) PrintCards() (bothPlayersCardsInHand string) {
	for i := 0; i < 2; i++ {
		bothPlayersCardsInHand += g.players[i].PrintCards() + "\n"

		// fmt.Println()

	}
	return bothPlayersCardsInHand
}
