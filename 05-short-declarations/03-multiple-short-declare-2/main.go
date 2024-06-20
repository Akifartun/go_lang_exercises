// Akif Artun
// Go Programming Exercises
// 03-multiple-short-declare-2

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Multiple Short Declare #2
//
//  1. Declare two variables using multiple short declaration
//
//  2. `a` variable's value should be 42
//  3. `c` variable's value should be "good"
//
// EXPECTED OUTPUT
//  42 good
// ---------------------------------------------------------

func main() {
	a, c := 42, "good"
	fmt.Println(a, c)
}
