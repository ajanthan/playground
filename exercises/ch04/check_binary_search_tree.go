package ch04

import "playground/ds/tree"

func IsBST(root *tree.BTNode) bool {
	if root == nil {
		return true
	}
	if !checkLeftDescendents(root.Left, root) || !checkRightDescendents(root.Right, root) {
		return false
	}
	return IsBST(root.Left) || IsBST(root.Right)
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
