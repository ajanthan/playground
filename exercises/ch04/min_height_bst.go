package ch04

import (
	"playground/ds/tree"
)

func MinHeightBST(in []int) *tree.BTNode {
	if len(in) == 1 {
		return &tree.BTNode{Data: in[0]}
	} else if len(in) == 0 {
		return nil
	}
	root := &tree.BTNode{}
	mid := len(in) / 2
	root.Data = in[mid]
	left := MinHeightBST(in[:mid])
	root.Left = left
	right := MinHeightBST(in[mid+1:])
	root.Right = right
	return root
}
