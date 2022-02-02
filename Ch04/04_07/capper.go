package main

import (
	"fmt"
	"os"
)

// Capper implements io.Writer and turns everything to uppercase
type Capper struct {
	// TODO
}

func main() {
	c := &Capper{os.Stdout}
	fmt.Fprintln(c, "Hello there")
}
