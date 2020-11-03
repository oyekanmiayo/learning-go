package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Ayo"
	names[1] = "Yuwa"
	fmt.Println(names)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)
}