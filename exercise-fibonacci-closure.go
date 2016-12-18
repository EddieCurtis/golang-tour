package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	oldA := 0
	oldB := 1
	return func() int {
		newA := oldB
		oldB = newA + oldA
		oldA = newA
		return newA
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
