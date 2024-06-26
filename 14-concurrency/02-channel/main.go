// Akif Artun
// Go Programming Exercises
// 02-channel

package main

import (
	"fmt"
)

// Channels are a typed conduit through which
// you can send and receive values with the channel operator, <-.
// The data flows in the direction of the arrow.

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, -1, 3, 5, 4}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
}
