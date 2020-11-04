package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

var k = map[string]Vertex{
	"Bell Labs" : Vertex{
		40.68433, -74.39967,
	}, 
	"Google" : Vertex{
		37.42202, -122.08408,
	},
}

var j = map[string]Vertex{
	"Bell Labs" : {40.68433, -74.39967},
	"Google" : {37.42202, -122.08408},
}

func main()  {
	
	// nothing can be added to m now because it is nil, we must use make
	fmt.Println(m)
	
	m = make(map[string]Vertex)
	fmt.Println(m)

	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	fmt.Println(k)

	fmt.Println(j)
}