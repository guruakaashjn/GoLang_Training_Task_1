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

	gameObj1 := guru_game.NewGame(user1, user2)
	gameObj1.Play()
	fmt.Println("\nGame ended")
}
