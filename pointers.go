package main

import "fmt"

// This is known as "dereferencing" or "indirecting
func main() {
	i, j := 42, 1005

	p := &i // point to i
	fmt.Println(*p) // read i through p
	*p = 4 // set i through p
	fmt.Println(i)

	p = &j
	*p = *p  / 35
	fmt.Println(j) // set j through p
}