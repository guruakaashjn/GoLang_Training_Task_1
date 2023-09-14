package main

import "fmt"

func isPrime(num int) int {
	if num <= 1 {
		return 0
	}
	if num == 2 || num == 3 {
		return 1
	}
	if num%2 == 0 || num%3 == 0 {
		return 0
	}

	for i := 5; i*i <= num; i = i + 6 {
		if num%i == 0 || num%(i+2) == 0 {
			return 0
		}
	}
	return 1

}

func main() {
	var num int
	fmt.Println("Enter a new number: ")
	fmt.Scan(&num)

	if isPrime(num) == 0 {
		fmt.Println("Not Prime")
	} else {
		fmt.Println("Prime")
	}

}
