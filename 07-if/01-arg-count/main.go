// Akif Artun
// Go Programming Exercises
// 01-arg-count

package main

import (
	"fmt"
	"os"
)

// ---------------------------------------------------------
// EXERCISE: Arg Count
//
//  1. Get arguments from command-line.
//  2. Print the expected outputs below depending on the number
//     of arguments.
//
// EXPECTED OUTPUT
//  go run main.go
//    Give me args
//
//  go run main.go hello
//    There is one: "hello"
//
//  go run main.go hi there
//    There are two: "hi there"
//
//  go run main.go I wanna be a gopher
//    There are 5 arguments
// ---------------------------------------------------------

func main() {
	input := os.Args
	l_input := len(input) - 1

	if l_input == 0 {
		fmt.Println("Give me args")
	} else if l_input == 1 {
		fmt.Printf("There is one: %s", input[1])
	} else if l_input == 2 {
		fmt.Printf("There is two: %s %s", input[1], input[2])
	} else {
		fmt.Printf("There are %d arguments", l_input)
	}
}
