package ch04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBST(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5, 6, 7}
	root := MinHeightBST(in1)
	assert.True(t, IsBST(root))
	root.Left.Right.Data = 5
	assert.False(t, IsBST(root))

	in2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	root = MinHeightBST(in2)
	assert.True(t, IsBST(root))
	root.Right.Right.Left.Data = 7
	assert.False(t, IsBST(root))
}
