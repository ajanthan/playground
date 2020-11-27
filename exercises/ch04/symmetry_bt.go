package ch04

import (
	"playground/ds/tree"
)

func IsSymmetry(root *tree.BTNode) bool {
	if root == nil {
		return false
	}
	t1, t2 := root.Left, root.Right
	if t1 == nil || t2 == nil {
		return false
	}
	return isSymmetryHelper(t1, t2)
}

func isSymmetryHelper(t1 *tree.BTNode, t2 *tree.BTNode) bool {
	if t1 == nil {
		return t2 == nil
	} else if t2 == nil {
		return t1 == nil
	}
	if t1.Data != t2.Data {
		return false
	}
	return isSymmetryHelper(t1.Left, t2.Right) && isSymmetryHelper(t1.Right, t2.Left)
}

func Reverse(root *tree.BTNode) *tree.BTNode {
	if root == nil {
		return nil
	}
	if root.Left == nil && root.Right == nil {
		return root
	}
	temp := root.Left
	root.Left, root.Right = root.Right, temp
	Reverse(root.Left)
	Reverse(root.Right)
	return root

}

func isEqual(t1 *tree.BTNode, t2 *tree.BTNode) bool {
	if t1 == nil {
		return t2 == nil
	} else if t2 == nil {
		return t1 == nil
	}
	if t1.Data != t2.Data {
		return false
	}
	return isEqual(t1.Left, t2.Left) && isEqual(t1.Right, t2.Right)
}
