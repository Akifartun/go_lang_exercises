// Akif Artun
// Go Programming Exercises
// 02-multiple-short-declare

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multiple Short Declare
//
//  Declare two variables using multiple short declaration
//
// EXPECTED OUTPUT
//  14 true
// ---------------------------------------------------------

func main() {
	number, valid := 14, true
	fmt.Println(number, valid)
}
