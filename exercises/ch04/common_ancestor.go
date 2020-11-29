package ch04

import (
	"math"
	"playground/ds/tree"
)

const (
	FOUND = iota
	NOTFOUND
	MARKED
)

func FindCommonAncestor(root *tree.BTNode, x, y int) int {
	if root == nil {
		return math.MinInt64
	}
	r, sr := findCommonAncestorHelper(root, x, y)
	if sr == MARKED {
		return r
	}
	return math.MinInt64
}

func findCommonAncestorHelper(node *tree.BTNode, x, y int) (int, int) {
	if node == nil {
		return math.MinInt64, NOTFOUND
	}
	if node.Data.(int) == x || node.Data.(int) == y {
		return math.MaxInt64, FOUND
	} else if node.Data.(int) == y {
		return math.MaxInt64, FOUND
	}
	l, sl := findCommonAncestorHelper(node.Left, x, y)
	r, sr := findCommonAncestorHelper(node.Right, x, y)
	if sl == NOTFOUND && sr == NOTFOUND {
		return math.MinInt64, NOTFOUND
	}
	if sl == FOUND && sr == FOUND {
		return node.Data.(int), MARKED
	} else if sl == FOUND {
		return l, sl
	} else if sr == FOUND {
		return r, sr
	} else if sl == MARKED {
		return l, MARKED
	} else if sr == MARKED {
		return r, MARKED
	}
	return -1, -1
}
