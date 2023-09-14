package main

import "fmt"

func main() {
	var arr_len int
	fmt.Println("Enter the length of array: ")
	fmt.Scan(&arr_len)

	var arr = make([]int, arr_len)
	fmt.Println("Enter array elements: ")
	for i := 0; i < arr_len; i++ {
		fmt.Scan(&arr[i])
	}

	for i := 0; i < arr_len; i++ {
		for j := i + 1; j < arr_len; j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}

	fmt.Println("Array after sorting")

	for i := 0; i < arr_len; i++ {
		fmt.Print(arr[i], " ")
	}

}
