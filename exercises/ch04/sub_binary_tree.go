package ch04

import "playground/ds/tree"

func IsSubBT(t1 *tree.BTNode, t2 *tree.BTNode) bool {
	if t1 == nil {
		return false
	}
	if t1.Data == t2.Data {
		if IsEqual(t1, t2) {
			return true
		}
	}
	return IsSubBT(t1.Left, t2) || IsSubBT(t1.Right, t2)
}

func IsEqual(t1 *tree.BTNode, t2 *tree.BTNode) bool {
	if t1 == nil || t2 == nil {
		return t1 == t2
	}
	if t1.Data != t2.Data {
		return false
	}
	return IsEqual(t1.Left, t2.Left) && IsEqual(t1.Right, t2.Right)
}
