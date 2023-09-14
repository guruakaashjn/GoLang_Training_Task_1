//// J N Guru Akaash
//// GoLang Training Course Practice

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var turn_score int
	turn_score = 0
	var total_score int
	total_score = 0

	random_number := 0
	turns := 0

start_turn:

	turn_score = 0
	fmt.Printf("\n")
	fmt.Printf("TURN %d:", turns+1)
	fmt.Printf("\n\n")

	fmt.Printf("Welcome to the game of Pig!")
	fmt.Println()

	for {

		fmt.Printf("\nEnter 'r' to roll again, 'h' to hold.\n")
		var choice string
		fmt.Scan(&choice)
		//fmt.Println(choice)

		switch choice {
		case "r":
			random_number = rand.Intn(6-1) + 1
			//fmt.Print(random_number)

			fmt.Printf("You rolled: %d", random_number)

			if random_number == 1 {
				fmt.Printf("\nTurn over. No Score\n")
				turns++
				goto start_turn
			} else {
				turn_score += random_number
				fmt.Printf("\nYour turn score is %d and your total score is %d", turn_score, total_score)
				fmt.Printf("\nIf you hold, you will have %d points.", turn_score+total_score)
			}
		case "h":
			total_score += turn_score

			turns++
			fmt.Printf("\nYour turn score is %d and your total score is %d\n", turn_score, total_score)
			goto start_turn

		}
		if turn_score+total_score >= 20 {
			fmt.Printf("\nYou Win! You finished in %d turns!", turns+1)
			fmt.Printf("\n\nGame Over!")
			break
		}

	}

}
