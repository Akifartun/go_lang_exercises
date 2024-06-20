// Akif Artun
// Go Programming Exercises
// 05-short-discard

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Short Discard
//
// 	1. Short declare two bool variables
//     (use multiple short declaration syntax)
//
//  2. Initialize both variables to true
//
//  3. Change your declaration and
//     discard the 2nd variable's value
//     using the blank-identifier
//
//  4. Print only the 1st variable
//
// EXPECTED OUTPUT
//  true
// ---------------------------------------------------------

func main() {
	a, _ := true, true
	fmt.Println(a)
}
