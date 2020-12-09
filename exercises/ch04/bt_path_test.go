package ch04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindPathsWithWight(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5, 6, 7, 1, 3, 5, 6, 7, 8, 9, -1, 2,
		3, 4, 5, 6, 7, 8, 4, 5, 7, 3, 5, 7, 4, -3, 4, 5, 6, 2, -3,
		6, 7, 3, 2, 9, 4, 7, 1, 5, 8, 9, 4, 5, -4, 3}
	root := MinHeightBST(in1)
	paths1 := FindPathsWithWightV0(root, 3)
	paths2 := FindPathsWithWightV1(root, 3)
	assert.Equal(t, paths1, paths2)
	paths1 = FindPathsWithWightV0(root, 8)
	paths2 = FindPathsWithWightV1(root, 8)
	assert.Equal(t, paths1, paths2)
	paths1 = FindPathsWithWightV0(root, 12)
	paths2 = FindPathsWithWightV1(root, 12)
	assert.Equal(t, paths1, paths2)
	paths1 = FindPathsWithWightV0(root, 17)
	paths2 = FindPathsWithWightV1(root, 17)
	assert.Equal(t, paths1, paths2)
}
