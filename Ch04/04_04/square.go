package main

import (
	"fmt"
	"log"
)

// Square is a square
type Square struct {
	// TODO
}

// NewSquare returns a new square
func NewSquare(x int, y int, length int) (*Square, error) {
	// TODO
	return nil, nil
}

// Move moves the square
func (s *Square) Move(dx int, dy int) {
	// TODO
}

// Area returns the square are
func (s *Square) Area() int {
	// TODO
	return 0
}

func main() {
	s, err := NewSquare(1, 1, 10)
	if err != nil {
		log.Fatalf("ERROR: can't create square")
	}

	s.Move(2, 3)
	fmt.Printf("%+v\n", s)
	fmt.Println(s.Area())
}
