package main

import (
	"fmt"
	"math"
)

const delta = 0.00001

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if (x < 0) {
		return 0, ErrNegativeSqrt(x)
	}
	previous := 1.0
	next := 0.0
	for {
		next = previous - (math.Pow(previous, 2) - x) / (2 * previous)
		if (math.Abs(next - previous) < delta) {
			return next, nil
		}
		previous = next
	}
	return 0, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
