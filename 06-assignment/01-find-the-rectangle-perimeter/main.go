// Akif Artun
// Go Programming Exercises
// 01-find-the-rectangle-perimeter

package main

import "fmt"

// ---------------------------------------------------------
// EXERCISE: Find the Rectangle's Perimeter
//
//  1. Find the perimeter of a rectangle
//     Its width is  5
//     Its height is 6
//
//  2. Assign the result to the `perimeter` variable
//
//  3. Print the `perimeter` variable
//
// HINT
//  Formula = 2 times the width and height
//
// EXPECTED OUTPUT
//  22
//
// ---------------------------------------------------------

func main() {
	var (
		width  int
		height int
	)

	width = 5
	height = 6

	var perimeter = 2 * (width + height)
	fmt.Println(perimeter)
}
