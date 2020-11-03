package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11}
	
	s = s[1:4]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = s[:2]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = s[:3]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = s[1:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)

	s = s[:]
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}