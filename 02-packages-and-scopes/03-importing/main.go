// Akif Artun
// Go Programming Exercises
// 03-importing

package main

// ---------------------------------------------------------
// EXERCISE: Rename imports
//
//  1- Import fmt package three times with different names
//
//  2- Print a few messages using those imports
//
// EXPECTED OUTPUT
//  hello
//  hey
//  hi
// ---------------------------------------------------------

import (
	f "fmt"
	fm "fmt"
	fmt "fmt"
)

func main() {
	f.Println("hello")
	fm.Println("hey")
	fmt.Println("hi")
}
