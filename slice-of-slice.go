package main


import(
	"fmt"
	"strings"
)


func main(){
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}

	board[0][1] = "X"
	board[1][1] = "O"
	board[1][1] = "X"
	board[0][0] = "O"
	board[1][1] = "X"


	for i := 0; i < len(board); i++{
		fmt.Printf("%s\n", strings.Join(board[i], ""))
	}
}

