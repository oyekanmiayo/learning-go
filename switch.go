package main

import(
	"fmt"
	"runtime"
)

func main()  {
	fmt.Print("Go runs on")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println(" OS X.")
	case "linux":
		fmt.Println(" linux.")
	default:
		fmt.Printf(" %v.\n", os)
	}

	fmt.Println(byte('A'))
}