package main

import "fmt"


func main(){
	s := []int{1,2,3,4,5,6,7,8,9,0}
	s = s[5:6]
	fmt.Println(s)
	
	s = s[:4]
	fmt.Println(s)
	

	fmt.Println("kita bisa memotong array nya dengan slice")
}


