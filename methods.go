package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X,  Y float64
}

// Remember: a method is just a function with a receiver argument.
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	j :=  Vertex{3, 4}
	fmt.Println(j.Abs())
}