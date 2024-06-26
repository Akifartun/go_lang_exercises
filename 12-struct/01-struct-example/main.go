// Akif Artun
// Go Programming Exercises
// 01-struct-example

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X, v.Y)
}
