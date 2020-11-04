package main

import "fmt"

/*
A sice has both a length an

*/
func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Slice the slice to give it a zero length
	s = s[:0]
	printSlice(s)

	// Extend its length
	s = s[:4]
	printSlice(s)

	// Drop the first 2 values in the underlying array
	s = s[2:6]
	printSlice(s)

	s = s[1:]
	printSlice(s)
	
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}