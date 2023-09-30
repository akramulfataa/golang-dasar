package main

import "fmt"

func main(){
	m := make(map[string]int)

	m["satu"] = 43
	fmt.Println("the values", m["satu"])

	m["satu"] = 89
	fmt.Println("thevalues", m["satu"])

	delete(m, "satu")
	fmt.Println("the values", m["satu"])

	v, ok := m["satu"]
	fmt.Println("the values", v, "presents?", ok)

}

