package main


import "fmt"

var pow = []int{1,2,3,4,5,6,7,8,9}

func main(){
	for i, v := range pow {
		fmt.Printf("no index %d = %d\n", i, v)
	}
}


