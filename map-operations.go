package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// Test that a key is present with a two-value assignment:
	// If key is in m, ok is true. If not, ok is false.
	// If key is not in the map, then elem is the zero value for the map's element type.
	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}