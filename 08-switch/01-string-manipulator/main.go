// Akif Artun
// Go Programming Exercises
// 01-string-manipulator

package main

import (
	"fmt"
	"os"
	"strings"
)

// ---------------------------------------------------------
// STORY
//  You want to write a program that will manipulate a
//  given string to uppercase and lowercase.
//
// EXERCISE: String Manipulator
//
//  1. Get the operation as the first argument.
//
//  2. Get the string to be manipulated as the 2nd argument.
//
// HINT
//  You can find the manipulation functions in the strings
//  package Go documentation (ToLower, ToUpper).
//
// EXPECTED OUTPUT
//
//  go run main.go
//    [command] [string]
//
//    Available commands: lower and upper
//
//  go run main.go lower 'OMG!'
//    omg!
//
//  go run main.go upper 'omg!'
//    OMG!
//
//  go run main.go genius "mr. charles darwin"
//    Unknown command: "genius"
// ---------------------------------------------------------

func main() {
	input := os.Args

	if len(input) == 1 {
		fmt.Println("Give me command and string")
		return
	}

	command := input[1]
	message := input[2]

	switch command {
	case "lower":
		l_message := strings.ToLower(message)
		fmt.Println(l_message)
	case "upper":
		u_message := strings.ToUpper(message)
		fmt.Println(u_message)
	default:
		fmt.Printf("Unkown command: \"%s\"", command)
	}
}
