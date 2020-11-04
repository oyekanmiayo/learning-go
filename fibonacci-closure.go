package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	 i := 1
	 j := 0
	return func() int {
		v := j
		i, j = j + i, i
		return v
	}
}

func main() {

	// 0, 1, 1
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
