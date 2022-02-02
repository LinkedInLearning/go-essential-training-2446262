package main

import (
	"fmt"
	"log"
)

// Square is a square
type Square struct {
	X      int
	Y      int
	Length int
}

// NewSquare returns a new square
func NewSquare(x int, y int, length int) (*Square, error) {
	if length <= 0 {
		return nil, fmt.Errorf("length must be > 0")
	}

	s := Square{
		X:      x,
		Y:      y,
		Length: length,
	}

	return &s, nil
}

// Move moves the square
func (s *Square) Move(dx int, dy int) {
	s.X += dx
	s.Y += dy
}

// Area returns the square are
func (s Square) Area() int {
	return s.Length * s.Length
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
