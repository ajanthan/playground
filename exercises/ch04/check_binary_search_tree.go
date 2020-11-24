package ch04

import (
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
