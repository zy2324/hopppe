package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func fib(n int) int {
	if n < 2 {
		return n
	}

	return fib(n-2) + fib(n-1)

	cur, next := 0, 1
	for i := 0; i < n; i++ {
		cur, next = next, cur+next
	}
	return cur
}
