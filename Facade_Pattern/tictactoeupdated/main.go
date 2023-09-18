package main

import (
	"fmt"
	"tictactoeupdated/game"
)

func main() {
	fmt.Println("Game started")
	var user1 string
	var user2 string

	fmt.Println("Enter User 1 Name: ")
	fmt.Scan(&user1)
	fmt.Println("Enter User 2 Name: ")
	fmt.Scan(&user2)
	g1 := game.NewGame(user1, user2)

	g1.Play()
	fmt.Println("Game ended")

}
