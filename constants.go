package main

import "fmt"

func main(){
	const name = "contoh"
	// const lebih dari satu 
	const (
		firstname = "akramul"
		lastname = "fata"
	)
	fmt.Println("nama depan", firstname)
	fmt.Println("nama belakang", lastname)

	// tidak bisa di dibuat dengan syntax := 
	// const age := 10
	
	fmt.Println(name)
}


