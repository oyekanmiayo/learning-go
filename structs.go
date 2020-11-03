package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

/*
Struct fields can be accessed through a struct pointer.

To access the field X of a struct when we have the struct pointer p we could write (*p).X. 
However, that notation is cumbersome, so the language permits us instead to write just p.X, without the explicit dereference.
*/
func main(){
	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v)

	p := &v
	p.Y = 1e3
	fmt.Println(v)
}