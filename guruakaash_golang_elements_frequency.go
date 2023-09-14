package main

import "fmt"

func countOccurances(arr []int, arr_len int) {
	var elem_freq = make([]int, 101)
	for i := 0; i < arr_len; i++ {
		elem_freq[arr[i]]++
	}

	for i := 0; i < arr_len; i++ {
		if elem_freq[arr[i]] != -1 {
			fmt.Printf("\n%d occured %d times", arr[i], elem_freq[arr[i]])
			elem_freq[arr[i]] = -1
		}

	}
}

func main() {
	var arr_len int
	fmt.Println("Enter length of array: ")
	fmt.Scan(&arr_len)

	var arr = make([]int, arr_len)
	fmt.Println("Enter elements of array: ")
	for i := 0; i < arr_len; i++ {
		fmt.Scan(&arr[i])
	}

	countOccurances(arr, arr_len)

}
