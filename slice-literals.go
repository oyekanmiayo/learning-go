package main

import "fmt"

/*
A slice literal is like an array literal without the length.

This is an array literal:
[3]bool{true, true, false}

And this creates the same array as above, then builds a slice that references it:
[]bool{true, true, false}
*/
func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{20, false},
		{78, false},
		{90, true},
		{b: false},
	}
	fmt.Println(s)

	var j []int
	fmt.Println(j, len(j), cap(j))

	a := make([]int, 5)
	printSlice("a", a)
}