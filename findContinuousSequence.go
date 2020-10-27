package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

// 一个数的所有求和
func findCoun(target int) [][]int {
	var arr []int
	for i := 1; i <= target; i++ {
		arr[i] = i
	}

	sum := 0
	var res [][]int

	for j := 0; j < len(arr); j++ {
		tmp := append(tmp, arr[j])
		sum += arr[i]
		if sum == target {
			res = append(res, tmp)
			sum := 0
		}
	}

	return res
}
