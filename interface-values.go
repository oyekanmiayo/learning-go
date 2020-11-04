package main

import "fmt"

type Vehicle interface {
	Move()
}

type Car struct {
	MovementAction string
}

func (c Car) Move() {
	fmt.Println(c.MovementAction)
}

type Plane struct {
	MovementAction string
}

func (p *Plane) Move() {
	fmt.Println(p.MovementAction)
}

func main() {
	var v Vehicle = Car{"Drive"}
	v.Move()
	describe(v)

	v = &Plane{"Fly"}
	v.Move()
	describe(v)

}

func describe(v Vehicle){
	fmt.Printf("(%v, %T)\n", v, v)
}