package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func deleteNode(head *ListNode, val int) *ListNode {
	for head != nil {
		if head.Var == val {
			head = head.Next
			break
		}
	}
	return head


	// 题解
	if head.Val == val {
		return head.Next
	}

	pre := head
	for head.Next.Val!=val{
		head=head.Next
	}

	head.Next=head.Next.Next
	return pre
}