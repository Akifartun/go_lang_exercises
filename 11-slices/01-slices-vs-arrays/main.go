// Akif Artun
// Go Programming Exercises
// 01-slices-vs-arrays

package main

import "fmt"

func main() {
	{
		var nums [5]int
		fmt.Printf("nums array: %#v\n", nums)
	}

	{
		var nums []int
		fmt.Printf("nums slice: %#v\n", nums)

		fmt.Printf("len(nums) : %d\n", len(nums))

		// fmt.Printf("nums[0]: %d\n", nums[0])
		// fmt.Printf("nums[1]: %d\n", nums[1])
	}
}
