package ch04

import (
	"math"
	"playground/ds/tree"
	"sort"
)

func IsBSTV1(root *tree.BTNode) bool {
	if root == nil {
		return true
	}
	if !checkLeftDescendents(root.Left, root) || !checkRightDescendents(root.Right, root) {
		return false
	}
	return IsBSTV1(root.Left) || IsBSTV1(root.Right)
}
func checkLeftDescendents(d *tree.BTNode, root *tree.BTNode) bool {
	if d == nil || root == nil {
		return true
	}
	if d.Data.(int) > root.Data.(int) {
		return false
	}
	if !checkLeftDescendents(d.Left, root) || !checkLeftDescendents(d.Right, root) {
		return false
	}
	return true
}
func checkRightDescendents(d *tree.BTNode, root *tree.BTNode) bool {
	if d == nil || root == nil {
		return true
	}
	if d.Data.(int) < root.Data.(int) {
		return false
	}
	if !checkRightDescendents(d.Left, root) || !checkRightDescendents(d.Right, root) {
		return false
	}
	return true
}

func BSTtoSlice(root *tree.BTNode) []int {
	s := make([]int, 0)
	return getRoot(root, s)
}

func getRoot(root *tree.BTNode, s []int) []int {
	if root.Left == nil {
		s = append(s, root.Data.(int))
		return s
	} else if root.Right == nil {
		s = append(s, root.Data.(int))
		return s
	}
	s = getRoot(root.Left, s)
	s = append(s, root.Data.(int))
	s = getRoot(root.Right, s)
	return s
}

func IsBSTV2(root *tree.BTNode) bool {
	s := BSTtoSlice(root)
	return sort.IntsAreSorted(s)
}

func isBSTV3(root *tree.BTNode, prevData int) (bool, int) {
	if root.Left == nil || root.Right == nil {
		if prevData == math.MinInt64 {
			return true, root.Data.(int)
		}
		return root.Data.(int) > prevData, root.Data.(int)
	}
	isOrdered, data := isBSTV3(root.Left, prevData)
	if !isOrdered {
		return false, data
	}
	isOrdered = root.Data.(int) > data
	prevData = root.Data.(int)
	if !isOrdered {
		return false, prevData
	}
	isOrdered, data = isBSTV3(root.Right, prevData)
	if !isOrdered {
		return false, data
	}
	return true, data
}
func IsBSTV3(root *tree.BTNode) bool {
	isBST, _ := isBSTV3(root, math.MinInt64)
	return isBST
}
