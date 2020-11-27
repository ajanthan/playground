package ch04

import (
	"github.com/stretchr/testify/assert"
	"math"
	"playground/ds/tree"
	"testing"
)

func TestFindNextNode(t *testing.T) {
	node1 := &tree.BBTNode{Data: 1}
	node2 := &tree.BBTNode{Data: 2}
	node1.Parent = node2
	node2.Left = node1

	node3 := &tree.BBTNode{Data: 3}
	node4 := &tree.BBTNode{Data: 4}
	node3.Parent = node4
	node4.Left = node3
	node4.Parent = node2

	node2.Right = node4

	node5 := &tree.BBTNode{Data: 5}

	node2.Parent = node5
	node5.Left = node2

	node6 := &tree.BBTNode{Data: 6}
	node5.Right = node6
	node6.Parent = node5

	node7 := &tree.BBTNode{Data: 7}
	node8 := &tree.BBTNode{Data: 8}
	node7.Parent = node8
	node8.Left = node7
	node8.Parent = node6

	node6.Right = node8

	node9 := &tree.BBTNode{Data: 9}
	node9.Left = node5
	node5.Parent = node9

	node10 := &tree.BBTNode{Data: 10}
	node10.Parent = node9
	node9.Right = node10

	i := FindNextNode(node2)
	assert.Equal(t, 3, i)

	i = FindNextNode(node4)
	assert.Equal(t, 5, i)

	i = FindNextNode(node8)
	assert.Equal(t, 9, i)

	i = FindNextNode(node10)
	assert.Equal(t, math.MinInt64, i)

}
