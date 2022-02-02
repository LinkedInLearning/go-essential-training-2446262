/*
An "even ended number" is a number who's first and last digit are the same.
You mission, should you choose to accept it, is to count how many "even ended
numbers" are there that are a multiplication of 4 digit numebrs.
*/

package main

import (
	"fmt"
)

func main() {
	/*
		n := 42
		s := fmt.Sprintf("%d", n)

		fmt.Printf("s = %q (type %T)\n", s, s)
	*/

	count := 0

	// for every pair of 4 digit numbers
	for a := 1000; a <= 9999; a++ {
		for b := a; b <= 9999; b++ { // don't count twice
			n := a * b

			// if a * b even ended
			s := fmt.Sprintf("%d", n)
			if s[0] == s[len(s)-1] {
				// count = count + 1
				count++
			}
		}
	}

	// print count
	fmt.Println(count)
}
