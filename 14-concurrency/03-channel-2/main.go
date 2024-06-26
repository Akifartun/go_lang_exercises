// Akif Artun
// Go Programming Exercises
// 03-channel-2

package main

import (
	"fmt"
)

// A sender can close a channel to indicate that
// no more values will be sent.
// Receivers can test whether a channel has been closed by
// assigning a second parameter to the receive expression: after

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10) // Create a buffered channel with capacity 10
	go fibonacci(cap(c), c) // Start a goroutine to generate Fibonacci numbers
	for i := range c {      // Receive values from the channel until it's closed
		fmt.Println(i)
	}
}
