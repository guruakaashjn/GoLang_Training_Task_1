//// J N Guru Akaash
//// GoLang Training Course Practice

package main

import "fmt"

func main() {
	var arr_len int

	fmt.Println("Enter array length: ")
	fmt.Scan(&arr_len)
	var Arr = make([]int, arr_len)

	fmt.Println("Enter array elements: ")
	for i := 0; i < arr_len; i++ {
		fmt.Scan(&Arr[i])
	}

	var largest int
	var second_largest int

	largest = Arr[0]
	second_largest = -1

	for i := 1; i < arr_len; i++ {
		if Arr[i] > largest {
			second_largest = largest
			largest = Arr[i]
		} else if Arr[i] < largest {
			if Arr[i] > second_largest {
				second_largest = Arr[i]
			}
		}
	}

	if second_largest == -1 {
		fmt.Println("No second largest number")
	} else {
		fmt.Printf("%d is your second largest number", second_largest)
	}

}
