package main

import "fmt"


func main() {
	fmt.Println("defer loop")
	i := 1
	for ; i < 10; i++ {
		defer fmt.Println("defer di dalam loop")
	}

	fmt.Println("done")

}

