package main

import "fmt"

func sumOfFib(n int) int {
	var sum int = 0
	if n == 1 {
		fmt.Println(1)
	} else if n == 2 {
		fmt.Println(2)
	}
	a := 1
	b := 1
	sum = a + b
	for n = n - 2; n != 0; n-- {
		temp := a + b
		a = b
		b = temp
		sum += temp

	}
	return sum
}

func main() {
	var n int
	fmt.Println("Enter number n: ")
	fmt.Scan(&n)

	var sum int = sumOfFib(n)
	fmt.Printf("Sum of fibonacci series upto n: %d", sum)
}
