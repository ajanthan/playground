package ch04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsSubBT(t *testing.T) {
	/*
	          4
	      2       6
	   1     3  5    7
	*/
	in1 := []int{1, 2, 3, 4, 5, 6, 7}
	t1 := MinHeightBST(in1)
	/*
	       2
	   1       3
	*/
	in2 := []int{1, 2, 3}
	t2 := MinHeightBST(in2)
	assert.True(t, IsSubBT(t1, t2))
	/*
	       2
	   1       4
	*/
	in3 := []int{1, 2, 4}
	t3 := MinHeightBST(in3)
	assert.False(t, IsSubBT(t1, t3))

	in4 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	t4 := MinHeightBST(in4)
	assert.True(t, IsSubBT(t4, t2))
}
