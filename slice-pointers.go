package main

import "fmt"


func main(){
	
	// declarasi array string
	names := [4]string{
		"fata",
		"aidil",
		"alam",
		"deby",
	}

	fmt.Println(names)
	a := names[0:2]
	b := names[2:3]
	fmt.Println(a, b)

	b[0] = "nurul"
	fmt.Println(a, b)
	fmt.Println(names)

}

