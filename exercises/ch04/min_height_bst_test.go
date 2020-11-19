package ch04

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestMinHeightBST(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	in2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16}
	in3 := []int{1, 3, 4, 8, 9, 10, 13, 17, 19, 20, 21, 22, 33, 44, 55, 100}
	root := MinHeightBST(in1)
	assert.Equal(t, 8, root.Data)

	root = MinHeightBST(in2)
	assert.Equal(t, 9, root.Data)

	root = MinHeightBST(in3)
	assert.Equal(t, 19, root.Data)
}
