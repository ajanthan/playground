package ch04

import (
	"math"
	"playground/ds/tree"
)

func FindNextNode(node *tree.BBTNode) int {
	if node == nil {
		return math.MinInt64
	}
	if node.Right != nil {
		return findNextNode(node.Right)
	}
	currentNode := node
	parentNode := node.Parent

	for parentNode != nil && parentNode.Left != currentNode {
		currentNode = currentNode.Parent
		parentNode = currentNode.Parent
	}
	if parentNode == nil {
		return math.MinInt64
	}
	return parentNode.Data.(int)

}
func findNextNode(node *tree.BBTNode) int {
	if node.Left == nil {
		return node.Data.(int)
	}
	return findNextNode(node.Left)

}
