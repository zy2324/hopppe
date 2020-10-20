package main

import (
	"fmt"
)

func main() {
	fmt.Println("vim-go")
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var pre, min int

func minDiffInBST(root *TreeNode) int {
	pre = -1
	min = math.MaxInt64
	minOrder(root)
	return min
}

func minOrder(root *TreeNode) {
	if root == nil {
		return 
	}
	minOrder(root.Left)
	if pre != -1 {
		if root.Val - pre < min {
			min = root.Val - pre
		}
	}
	pre = root.Val
	minOrder(root.Right)
}