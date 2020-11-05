package main

import "fmt"

func do(i interface{}){
	switch i.(type) {
	case int:
		fmt.Println("int type")
	case string:
		fmt.Println("string type")
	default:
		fmt.Printf("unknown type %T\n", i)
	}
}

func main()  {
	do(12)
	do("Ayo")
	do(4 + 5i)
	do(false)
}