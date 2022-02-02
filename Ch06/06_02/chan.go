// channels
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	//	ch <- 353 // send

	go func() {
		// Send number of the channel
		ch <- 353
	}()

	val := <-ch // receive
	fmt.Printf("got %d\n", val)

	fmt.Println("-----")

	// Send multiple
	const count = 3
	go func() {
		for i := 0; i < count; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for i := 0; i < count; i++ {
		val := <-ch
		fmt.Printf("received %d\n", val)
	}

	fmt.Println("-----")

	// close to signal we're done
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Printf("sending %d\n", i)
			ch <- i
			time.Sleep(time.Second)
		}
		close(ch)
	}()

	for i := range ch {
		fmt.Printf("received %d\n", i)
	}

	ch2 := make(chan int, 1) // buffered channel
	ch2 <- 19
	val2 := <-ch2
	fmt.Println(val2)
}
