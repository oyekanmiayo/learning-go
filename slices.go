package main

import "fmt"

/*
Slices are **references** to arrays.

A slice does not store any data, it just describes a section of an underlying array.

This selects a half-open range which includes the first element, but excludes the last one.

To include the last value, the second value in the range should be last index + 1
*/
func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	var pSlice []int = primes[2:6]
	fmt.Println(pSlice)

	names := [5]string{
		"Yuwa",
		"Jess",
		"Shay",
		"Chi",
		"Vee",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[2:5]
	fmt.Println(a, b)

	b[0] = "Seun"
	fmt.Println(a, b)
	fmt.Println(names)
}