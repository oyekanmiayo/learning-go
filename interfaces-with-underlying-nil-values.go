package main

import (
	"fmt"
	"math"
)
/*
If the interface holds a nil concrete value, any method is called with a nil receiver
If a method is called with a nil receiver, Go allows us to handle it gracefully

Also, if an interface value holds a nil concrete value, the interface value itself non-nil
*/

type Area interface {
	FindArea() float64
}

type  Circle struct  {
	radius float64
}

func (c *Circle) FindArea() float64 {
	if c == nil {
		fmt.Println("Circle object is nil")
		return 0
	}

	return math.Pi * c.radius * c.radius
}

func main()  {
	var a Area

	var c *Circle 
	a = c // c is nil
	fmt.Println(a.FindArea())

	/*
	var c Circle 
	if c == nil {}  keeps failing
	*/

	// interface a is not nil because it contains a nil Circle value
	if a == nil {
		fmt.Println("a is nil")
	}
}

func describe(a Area) {
	fmt.Printf("(%v, %T)\n", a, a)
}