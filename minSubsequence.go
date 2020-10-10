package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}


func minSubsequence(nums []int) []int {
	sort.Ints(nums)
	i, j := 0, len(nums)

	var res [][]int
	for i<j {
		left += nums[i]
		right += nums[i:j]

		if sum(left) >= sum(right) {
			var tmp []int
			tmp = append(tmp, left)
		}
	}
}