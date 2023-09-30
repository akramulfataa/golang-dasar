package main 

import "fmt"

func main(){
	angka := [5]int{1,2,2,2,2}
	
	var slice []int = angka[1:3]
	fmt.Println(slice)

}

