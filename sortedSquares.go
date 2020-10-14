package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}


func sortedSquares(A []int) []int {
	for i:=0; i<len(A); i++ {
		tmp := A[i]
		A[i] = tmp * tmp
	}
	return sort.Int(A)
}