package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func pivotIndex(nums []int) int {
	var num_left, num_right []int
	var res int = -1

	i, j := 0, len(nums)
	for i < j {
		num_left = append(num_left, nums[i])
		num_right = append(num_right, nums[j])

		if sum(num_left) == sum(num_right) {
			res = i
			break
		} else {
			i ++

		}

	}

	// 题解
	sum := 0
	for _, u := range nums {
		sum += u
	}

	left_sum := 0
	for i:=0; i<len(nums);i++{
		if left_sum == (sum-left_sum-nums[i]) {
			return i
		}
		left_sum += nums[i]
	}
	return -1
}

