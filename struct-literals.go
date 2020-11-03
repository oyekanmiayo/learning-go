package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

/*
Struct literals - all the different ways to construct structs
*/
var (
	v1 = Vertex{1, 2} // has type Vertex
	v2 = Vertex{X: 30} // Y : 0 is implicit
	v3 = Vertex{} // X : 0 and Y : 0 are implicit
	p = &Vertex{4, 5} // has type *Vertex
)
func main() {
	fmt.Println(v1, v2, v3, p)
}