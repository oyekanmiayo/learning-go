package main

import "fmt"

type Person struct{
	Name string
	Age int
}

func ChangeAge(p *Person, newAge int) {
	p.Age = newAge
}

func ChangeName(p *Person, newName string){
	p.Name = newName
}

func main(){
	p := Person{"Ayomide", 15}
	fmt.Println(p)

	ChangeAge(&p, 20)
	fmt.Println(p)

	ChangeName(&p, "Ayo")
	fmt.Println(p)
}