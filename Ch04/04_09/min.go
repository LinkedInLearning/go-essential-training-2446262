package main

import (
	"fmt"
)

type Ordered interface {
	int | float64 | string
}

func min[T Ordered](values []T) (T, error) {
	if len(values) == 0 {
		var zero T
		return zero, fmt.Errorf("min of empty slice")
	}

	m := values[0]
	for _, v := range values[1:] {
		if v < m {
			m = v
		}
	}
	return m, nil
}

func main() {
	fmt.Println(min([]float64{2, 1, 3}))
	fmt.Println(min([]string{"B", "A", "C"}))
}
