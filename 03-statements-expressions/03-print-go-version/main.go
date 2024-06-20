// Akif Artun
// Go Programming Exercises
// 03-print-go-version

package main

import (
	"fmt"
	"runtime"
)

// ---------------------------------------------------------
// EXERCISE: Print the Go Version
//
//  1. Look at the runtime package documentation
//  2. Find the func that returns the Go version
//  3. Print the Go version by calling that func
//
// EXPECTED OUTPUT
//  "go1.22.4"
// ---------------------------------------------------------

func main() {
	fmt.Println(runtime.Version())
}
