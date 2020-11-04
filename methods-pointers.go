package main

import "fmt"


/*
You can declare methods with pointer receivers.

This means the receiver type has the literal syntax *T for some type T. (Also, T cannot itself be a pointer such as *int.)

For example, the ChangeAge and ChangeName methods here are defined on *Person.

Methods with pointer receivers can modify the value to which the receiver points. This means that the values in the memory 
addresses are modified. 

Since methods often need to modify their receiver, pointer receivers are more common than value receivers.
*/

type Person struct{
	Name string
	Age int
}

func (p *Person) ChangeAge(newAge int) {
	p.Age = newAge
}

func (p *Person) ChangeName(newName string){
	p.Name = newName
}

func main()  {
	p := Person{"Ayomide", 45}
	fmt.Println(p)

	p.ChangeAge(30)
	fmt.Println(p)

	p.ChangeName("Ayo")
	fmt.Println(p)
}