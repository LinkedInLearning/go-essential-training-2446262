// Basic function definition
package main

import (
	"fmt"
)

func main() {
	val := add(1, 2)
	fmt.Println(val)

	div, mod := divmod(7, 2)
	fmt.Printf("div=%d, mod=%d\n", div, mod)
}

// add adds a to b
func add(a int, b int) int {
	return a + b
}

// divmod returns quotient and reminder
func divmod(a int, b int) (int, int) {
	return a / b, a % b
}
