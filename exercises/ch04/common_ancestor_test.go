package ch04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindCommonAncestor(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5, 6, 7}
	root := MinHeightBST(in1)
	i := FindCommonAncestor(root, 1, 7)
	assert.Equal(t, 4, i)

	i = FindCommonAncestor(root, 1, 3)
	assert.Equal(t, 2, i)
	i = FindCommonAncestor(root, 5, 7)
	assert.Equal(t, 6, i)
}
