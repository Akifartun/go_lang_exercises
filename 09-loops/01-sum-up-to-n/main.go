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
// EXERCISE: Sum up to N
//
//  1. Get two numbers from the command-line: min and max
//  2. Convert them to integers (using Atoi)
//  3. By using a loop, sum the numbers between min and max
//
// RESTRICTIONS
//  1. Be sure to handle the errors. So, if a user doesn't
//     pass enough arguments or she passes non-numerics,
//     then warn the user and exit from the program.
//
//  2. Also, check that, min < max.
//
// HINT
//  For converting the numbers, you can use `strconv.Atoi`.
//
// EXPECTED OUTPUT
//  Let's suppose that the user runs it like this:
//
//    go run main.go 1 10
//
//  Then it should print:
//
//    1 + 2 + 3 + 4 + 5 + 6 + 7 + 8 + 9 + 10 = 55
// ---------------------------------------------------------

func main() {
	input := os.Args
	if len(input) < 3 {
		fmt.Println("Give me 'min' and 'max' values.")
		return
	}

	min_value, err_min := strconv.Atoi(input[1])
	max_value, err_max := strconv.Atoi(input[2])

	if err_min != nil || err_max != nil {
		fmt.Println("Please give me numbers.")
		return
	}

	if max_value < min_value {
		fmt.Println("Max value is not greater than min value")
		return
	}

	sum := min_value
	fmt.Printf("%d ", min_value)
	for i := min_value + 1; i <= max_value; i++ {
		sum += i
		fmt.Printf("+ %d ", i)
	}

	fmt.Printf("= %d", sum)

}
