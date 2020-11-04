package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main()  {
	for  i, v :=  range pow  {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// To skip index
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}

	// To skip value
	for i, _ := range pow {
		fmt.Printf("%d\n", i)
	}
}