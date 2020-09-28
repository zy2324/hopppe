package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

func removeDuplicateNodes(head *ListNode) *ListNode {
	tmp := make(map[int]int)
	var res *ListNode
	for head != nil {
		if _, ok := tmp[head.Val]; ok {
			head = head.Next
		}
		res
	}
	// 差点火候
}
