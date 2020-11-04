package main

import "fmt"

func main() {
	a := []string{"Jam", "Butter"}
	b := []string{"Mayo", "Peanut-Butter"}
	// To append slices to each other use ... to expand second argument to
	// a list of arguments
	a = append(a, b...)
	printSlice(a)

	var s []int
	printSlice(s)

	// append works on nil slices.
	s = append(s, 0)
	printSlice(s)

	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)

	// We can add more than one element at a time.
	s = append(s, 2, 3, 4)
	printSlice(s)
}

func printSlice(s interface{}) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}