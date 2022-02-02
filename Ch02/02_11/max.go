// Calculate maximal value in a slice
package main

import (
	"fmt"
)

func main() {
	nums := []int{16, 8, 42, 4, 23, 15}
	// fmt.Println(nums)

	max := nums[0] // Initialize max with first value
	// [1:] skips the first value
	for _, value := range nums[1:] {
		if value > max {
			max = value
		}
	}

	fmt.Println(max)
}
