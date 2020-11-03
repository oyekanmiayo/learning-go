package main

import "fmt"

func main(){
	sum := 0
	for i := 1; i < 10; i++ {
		sum += 1
	}
	fmt.Println(sum)

	sum_two := 1
	for ; sum_two <  1000; {
		sum_two += sum_two
	}
	fmt.Println(sum_two)

	// Similar to `while` in other languagges
	sum_three := 1
	for sum_three < 1000 {
		sum_three += sum_three
	}
	fmt.Println(sum_three)
}