package main

import "LeetCodeGo/base"

var result = 0

/**
 * pseudoPalindromicPaths
 * 给定一棵二叉树,每个节点的值为 1 到 9
 * 二叉树中的一条路径是伪回文的需要满足:
 *  路径经过的所有节点值的排列中,存在一个回文序列
 * 请返回从根到叶子节点的所有路径中伪回文路径的数目
 *
 * 给定二叉树的节点数目在范围 [1, 10^5] 内
 * 1 <= Node.val <= 9
 *
 * @author Li.zongjie
 * @version 1.0
 * @date 2023/11/25
 */
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
