package main

import "fmt"

type Vehicle interface {
	Move()
}

type Car struct {
	MovementType string
}

// This method means type Car implements the interface Vehicle,
// but we don't need to explicitly declare that it does so.
func (c Car) Move() {
	fmt.Println(c.MovementType)
}

func main() {
	var v Vehicle = Car{"Accelerate"}
	v.Move()
}
