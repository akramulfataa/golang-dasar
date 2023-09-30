package main

import "fmt"

type Vertex struct{
	Lat, Long float64
}

var m = map[string]Vertex{
	"google" : Vertex{
		111111,33333,
	},
	"tubit" : Vertex{
		3333,4444,
	},
}

func main(){
	fmt.Println(m)
}



