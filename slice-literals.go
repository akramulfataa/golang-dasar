package main

import "fmt"


func main(){
	SliceLiteral := []int {1,2,3,4,5,5,}
	fmt.Println(SliceLiteral)

	// bisa pakek satu struct kosong
	s := []struct{
		name string
		boll bool
	}{
		{"akramul", true},
		{"fata", false},

	}
	fmt.Println(s)

}


