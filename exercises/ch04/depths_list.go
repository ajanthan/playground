package ch04

import (
	"playground/ds/tree"
)

func ListOfDepths(root *tree.BTNode) map[int][]*tree.BTNode {
	list := make(map[int][]*tree.BTNode)
	listOfDepths(root, list, 0)
	return list
}

func listOfDepths(root *tree.BTNode, list map[int][]*tree.BTNode, level int) {
	if root == nil {
		return
	}
	list[level] = append(list[level], root)
	listOfDepths(root.Left, list, level+1)
	listOfDepths(root.Right, list, level+1)
}
