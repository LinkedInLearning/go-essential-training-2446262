/*
An "even ended number" is a number whose first and last digit are the same.

You mission, should you choose to accept it, is to count how many "even ended numbers" are
there that are a multiplication of two 4 digit numbers.
*/

package main

import (
	"fmt"
)

func main() {
	n := 42
	s := fmt.Sprintf("%d", n)

	fmt.Printf("s = %q (type %T)\n", s, s)
}
