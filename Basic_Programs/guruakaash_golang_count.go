package main

import "fmt"

func main() {
	var arr_len int
	fmt.Println("Enter array length: ")
	fmt.Scan(&arr_len)

	var arr = make([]int, arr_len)
	fmt.Println("Enter the array elements: ")
	for i := 0; i < arr_len; i++ {
		fmt.Scan(&arr[i])
	}

	var count_even int
	var count_odd int
	var count_zero int

	for i := 0; i < arr_len; i++ {
		if arr[i] == 0 {
			count_zero++
		} else if arr[i]%2 == 0 {
			count_even++
		} else {
			count_odd++
		}
	}
	fmt.Printf("\nCount of even: %d", count_even)
	fmt.Printf("\nCount of odd: %d", count_odd)
	fmt.Printf("\nCount of zero: %d", count_zero)

}
