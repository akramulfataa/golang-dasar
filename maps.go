package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex


func main(){
	m = make(map[string]Vertex)

	m["Longtitude"] = Vertex{
		5.00000, 97.999999, 
	}
	fmt.Println(m["Longtitude"])
}


