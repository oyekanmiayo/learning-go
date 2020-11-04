package main

import (
	"fmt"
	"math"
)

type Area interface {
	FindArea() float64
}

type Circle struct {
	radius float64
}

func (c Circle) FindArea() float64 {
	return math.Pi * c.radius * c.radius
}

type Rectangle struct {
	length float64
	breadth float64
}

func (r Rectangle) FindArea() float64  {
	return r.length * r.breadth
}

func main(){
	var a Area
	c := Circle{6}
	r := Rectangle{2, 3}

	a = c
	fmt.Println(a.FindArea())

	a = r
	fmt.Println(a.FindArea())
}