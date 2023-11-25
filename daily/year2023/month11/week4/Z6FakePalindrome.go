package main

import "LeetCodeGo/base"

var result = 0

func pseudoPalindromicPaths(root *base.TreeNode) int {
	result = 0
	count(root, 0)
	return result
}

func count(root *base.TreeNode, value int) {
	if root == nil {
		return
	}
	value = value ^ (1 << root.Val)
	if root.Left == nil && root.Right == nil && value == (value&-value) {
		result++
	}
	count(root.Left, value)
	count(root.Right, value)
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
