package main

import "fmt"

type Vertex struct{
	Lat, Long float64
}

var m = map[string]Vertex{
	"landeng": {2838473,84858754},
	"tubit": {20745423,89592},
}

func main(){
	fmt.Println(m)
}

