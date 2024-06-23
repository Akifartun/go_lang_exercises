// Akif Artun
// Go Programming Exercises
// 03-leap-year

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ---------------------------------------------------------
// EXERCISE: Leap Year
//
//  Find out whether a given year is a leap year or not.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me a year number
//
//  go run main.go eighties
//    "eighties" is not a valid year.
//
//  go run main.go 2018
//    2018 is not a leap year.
//
//  go run main.go 2100
//    2100 is not a leap year.
//
//  go run main.go 2019
//    2019 is not a leap year.
//
//  go run main.go 2020
//    2020 is a leap year.
//
//  go run main.go 2024
//    2024 is a leap year.
// ---------------------------------------------------------

func main() {
	input := os.Args
	if len(input) == 1 {
		fmt.Println("Give me a year number")
		return
	}

	year, err := strconv.Atoi(input[1])
	if err != nil {
		fmt.Printf("\"%s\" is not valid year.", input[1])
		return
	}

	if year%4 == 0 && (year%100 != 0 || year%400 == 0) {
		fmt.Printf("\"%d\" is a leap year.", year)
	} else {
		fmt.Printf("\"%d\" is not a leap year.", year)
	}
}
