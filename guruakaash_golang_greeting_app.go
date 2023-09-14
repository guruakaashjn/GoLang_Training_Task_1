package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var time_input string
	var hh int
	var ss int

	fmt.Println("Enter time: ")
	fmt.Scan(&time_input)

	if idx := strings.IndexByte(time_input, ':'); idx >= 0 {
		h, err := strconv.Atoi(time_input[:idx])

		hh = h
		if err != nil {
			panic(err)
		}
	}

	// mm, err := strconv.Atoi(time_input[3:5])
	// if err != nil {
	// 	panic(err)
	// }

	if idx := strings.IndexByte(time_input, ':'); idx >= 0 {
		temp := time_input[idx+1:]

		if idxx := strings.IndexByte(temp, ':'); idxx >= 0 {
			s, err := strconv.Atoi(time_input[idx+1 : idx+3])
			ss = s
			if err != nil {
				panic(err)
			}
		}

	}

	var am_or_pm string

	var temp string
	temp = "a"

	if string(time_input[len(time_input)-2]) == temp {
		am_or_pm = "am"
	} else {
		am_or_pm = "pm"
	}

	if hh > 6 && hh < 11 && am_or_pm == "am" || hh == 6 && ss >= 01 && am_or_pm == "am" || hh == 11 && ss == 0 && am_or_pm == "am" {
		fmt.Println("Good morning!")
	} else if hh > 11 && am_or_pm == "am" || hh == 11 && ss >= 01 && am_or_pm == "am" || hh < 4 && am_or_pm == "pm" || hh == 4 && ss == 0 && am_or_pm == "pm" {
		fmt.Println("Good afternoon!")
	} else if hh > 4 && hh < 9 && am_or_pm == "pm" || hh == 4 && ss >= 01 && am_or_pm == "pm" || hh == 9 && ss == 0 && am_or_pm == "pm" {
		fmt.Println("Good evening!")
	} else if hh > 9 && am_or_pm == "pm" || hh == 9 && ss >= 01 && am_or_pm == "pm" || hh < 6 && am_or_pm == "am" || hh == 6 && ss == 0 && am_or_pm == "am" {
		fmt.Println("Good night!")
	}

}
