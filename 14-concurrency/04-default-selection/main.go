// Akif Artun
// Go Programming Exercises
// 04-default-selection

package main

import (
	"fmt"
	"time"
)

// A select blocks until one of its cases can run,
// then it executes that case.
// It chooses one at random if multiple are ready.

func main() {
	tick := time.Tick(500 * time.Millisecond)
	boom := time.After(1500 * time.Millisecond)
	for {
		select {
		case <-tick:
			fmt.Println("tick. ")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("    .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
