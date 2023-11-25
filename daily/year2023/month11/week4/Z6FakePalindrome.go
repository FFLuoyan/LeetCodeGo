package main

import "LeetCodeGo/base"

var result = 0

func pseudoPalindromicPaths(root *base.TreeNode) int {
	result = 0
	count(root, 0)
	return result
}

func count(root *base.TreeNode, value int) {
	value = value ^ (1 << root.Val)
	if root.Left == nil {
		if root.Right == nil {
			if value == (value & -value) {
				result++
			}
		} else {
			count(root.Right, value)
		}
	} else if root.Right == nil {
		count(root.Left, value)
	} else {
		count(root.Left, value)
		count(root.Right, value)
	}
}

func main() {
	// 1
	print(pseudoPalindromicPaths(
		&base.TreeNode{
			Val: 2,
			Left: &base.TreeNode{
				Val: 1,
				Left: &base.TreeNode{
					Val: 1,
				},
				Right: &base.TreeNode{
					Val: 3,
					Right: &base.TreeNode{
						Val: 1,
					},
				}},
			Right: &base.TreeNode{
				Val: 1,
			},
		}))
}
