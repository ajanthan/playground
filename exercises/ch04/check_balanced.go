package ch04

import (
	"playground/ds/tree"
)

func getHeight(root *tree.BTNode, level int) (bool, int) {
	if root == nil {
		return false, 0
	}
	if root.Left == nil && root.Right == nil {
		return true, 1
	}
	lHeight, rHeight := 0, 0
	lBal, rBal := false, false
	if root.Left != nil {
		lBal, lHeight = getHeight(root.Left, level+1)
	}
	if root.Right != nil {
		rBal, rHeight = getHeight(root.Right, level+1)
	}

	if !lBal || !rBal {
		return false, 0
	}
	diff := lHeight - rHeight
	if diff > 1 || diff < -1 {
		return false, 0
	}
	return true, level + 1
}

func CheckBalanced(root *tree.BTNode) bool {
	ret, _ := getHeight(root, 0)
	return ret

}
