// Akif Artun
// Go Programming Exercises
// 05-mutex

package main

import (
	"fmt"
	"sync"
	"time"
)

// Go's standard library provides mutual exclusion
// with sync.Mutex and its two methods:

type SafeCounter struct {
	mu sync.Mutex
	v  map[string]int
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	c.v[key]++
	c.mu.Unlock()
}

func (c *SafeCounter) Value(key string) int {
	c.mu.Lock()
	// Lock so only one goroutine at a time can access the map c.v.
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	i := 0
	for i < 1000 {
		go c.Inc("mutexes")

		i += 1
	}
	// If there is a "go" statement before "c.Inc("mutexes")",
	// there must be time.sleep to wait until finish loop
	// otherwise before the incremental done it will print sum number.

	// Second alternative is the remove "go" statement before "c.Inc("mutexes")"
	time.Sleep(time.Second)
	fmt.Println(c.Value("mutexes"))
}
