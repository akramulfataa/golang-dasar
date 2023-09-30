package main

import "fmt"

func main(){
	s := []int{1,2,3,4,5,6,7,8,9,10}
	// buat func untuk print slice nya
	printSlice(s)

	s = s[:0]
	printSlice(s)
	// ubah slice nya jadi 0

	s = s[0:4]
	printSlice(s)
	
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

