package ch04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckBalanced(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5, 6, 7}
	root := MinHeightBST(in1)
	assert.True(t, CheckBalanced(root))
	root.Right = nil
	assert.False(t, CheckBalanced(root))
	in2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	root = MinHeightBST(in2)
	assert.True(t, CheckBalanced(root))
	root.Right.Right = nil
	assert.False(t, CheckBalanced(root))
}
