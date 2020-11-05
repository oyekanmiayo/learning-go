package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

/*
Basically, Go's version of `toString()` in Java
*/
func (p Person) String() string {
	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
}

func main(){
	a := Person{"Ayo", 20}
	b := Person{"Yuwa", 20}
	fmt.Println(a, b)
	// Output without `String()`: {Ayo 20} {Yuwa 20}
	// Output with `String()`: Ayo (20 years) Yuwa (20 years)
}