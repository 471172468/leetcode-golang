package main

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	tn := []*TreeNode{root}
	for len(tn) != 0 {
		t := tn[0]
		tn = tn[1:]
		t.Left, t.Right = t.Right, t.Left	// 交换
		if t.Left != nil {
			tn = append(tn, t.Left)
		}
		if t.Right != nil {
			tn = append(tn, t.Right)
		}
	}
	return root
}
