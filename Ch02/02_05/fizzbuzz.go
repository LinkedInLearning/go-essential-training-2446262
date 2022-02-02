/*
Print the numbers between 1 to 20, however
- If the number is divisible by 3, print "fizz" instead
- If the number is divisible by 5, print "buzz" instead
- If the number is divisible by 3 and 5, print "fizz buzz" instead
- Otherwise print the number
*/

package main

import (
	"fmt"
)

func main() {
	/*
		fmt.Println(1 % 5)
		fmt.Println(7 % 5)
		fmt.Println(10 % 5)
	*/

	// for every number from 1 to 20
	for i := 1; i <= 20; i++ {
		if i%3 == 0 && i%5 == 0 {
			// if the number is divisible by 3 and 5 print fizz buzz
			fmt.Println("fizz buzz")
		} else if i%3 == 0 {
			// else if the number is divisible by 3 print fizz
			fmt.Println("fizz")
		} else if i%5 == 0 {
			// else if the number is divisible by 5 print buzz
			fmt.Println("buzz")
		} else {
			// else print the number
			fmt.Println(i)
		}
	}
}
