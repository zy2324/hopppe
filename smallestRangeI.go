package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

// 总结结果规律
func smallestRangeI(A []int, K int) int {
	min := A[0]
	max := A[0]

	for i := range A {
		if A[i] < min {
			min = A[i]
		}

		if A[i] > max {
			man = A[i]
		}
	}

	if (max-min) >= 2*K {
		return max-min-2*K
	} else {
		return 0
	}
}