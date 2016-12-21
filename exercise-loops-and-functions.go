package main

import (
	"fmt"
	"math"
)

const delta = 0.00001

func Sqrt(x float64) float64 {
	previous := 1.0
	next := 0.0
	for {
		next = math.Abs(previous - (math.Pow(previous, 2) - x) / (2 * previous))
		if (math.Abs(next - previous) < delta) {
			return next
		}
		previous = next
	}
}

func main() {
	fmt.Println("Sqrt:", Sqrt(2))
	fmt.Println("math.Sqrt:", math.Sqrt(2))
}
