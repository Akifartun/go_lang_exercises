// Akif Artun
// Go Programming Exercises
// 01-goroutine

package main

import (
	"fmt"
	"time"
)

// A goroutine is a lightweight thread managed by the Go runtime.

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("World!")
	say("Hello")
}
