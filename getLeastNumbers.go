package main

import (
	"fmt"
)

func main() {
	arr := []int{3, 2, 2, 4}
	fmt.Println(getLeastNumbers(arr, 2))
}

func getLeastNumbers(arr []int, k int) []int {
	return quickSort(arr[:])[:k]
}

func quickSort(arr []int) (res []int) {
	if len(arr) < 2 {
		return arr
	}

	pivot := arr[0]
	var left, middle, right []int

	for _, v := range arr {
		if v < pivot {
			left = append(left, v)
		} else if v == pivot {
			middle = append(middle, v)
		} else if v > pivot {
			right = append(right, v)
		}
	}

	return append(append(quickSort(left[:]), middle...), quickSort(right[:])...) // 递归排序
}
