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

	// g1.Play()
	// fmt.Println("Game ended")
	// g1.Play()
	// g1.Play()

	for i := 0; i < 1; {
		// startturn:

		fmt.Printf("\nTurn %d\n", g1.GetTurn())
		fmt.Printf("Enter cell number: ")
		var cellNumber uint
		fmt.Scan(&cellNumber)
		var flag, response, boardMarks = g1.PlayLogic(cellNumber)
		fmt.Println(boardMarks)
		if flag {
			i++
		}
		fmt.Println(response)
	}

	var flag, response, boardMarks = g1.PlayLogic(1)
	fmt.Println("F : ", flag)
	fmt.Println("Response: ", response)
	fmt.Println(boardMarks)
}
