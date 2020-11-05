package main
 
import (
    "bytes"
    "fmt"      
)

func main()  {
	countries := []byte("Australia Canada Japan Germany India")
    space := []byte{' '}
	splitExample := bytes.Split(countries, space)
	
	fmt.Println("Countries", string(countries))
	fmt.Println("Split", splitExample)

	nums := [4]byte{1, 2, 4, 5}
	fmt.Println(nums)
}