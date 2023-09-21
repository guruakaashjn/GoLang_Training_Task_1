package main

import (
	"blackjack/guru_game"
	"fmt"
)

func main() {
	fmt.Println("Main called....")
	var user1 string
	var user2 string

	fmt.Println("Enter User 1 Name: ")
	fmt.Scan(&user1)
	fmt.Println("Enter User 2 Name: ")
	fmt.Scan(&user2)

	gameObj1, initialTurn := guru_game.NewGame(user1, user2)
	fmt.Println(initialTurn)

	for i := 0; i < 1; {
		// fmt.Println(gameObj1.GetTurn())
		// gameObj1.PrintCards()
		fmt.Println("1. Draw a card")
		fmt.Println("2. Hold cards")
		var choice int
		fmt.Scan(&choice)

		flag, response, bothPlayersCardsInHand := gameObj1.Play(choice)
		fmt.Println(bothPlayersCardsInHand)
		if flag {
			i++
		}
		fmt.Println(response)

		// fmt.Println("Flag: ", flag)
		// fmt.Println("Response: ", response)
	}

	fmt.Println("\nGame ended")

	_, response, bothPlayersCardsInHand := gameObj1.Play(1)
	fmt.Println(bothPlayersCardsInHand)
	fmt.Println(response)

}
