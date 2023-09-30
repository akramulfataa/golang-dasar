package main 

import "fmt"

func cinta(x, y, z string) (string, string, string){
	return x, y, z
}

func main() {
	aku, cinta, kamu := cinta("aku", "cinta","kamu")
	fmt.Println(aku, cinta, kamu)
}

