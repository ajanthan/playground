package ch04

import (
	"github.com/stretchr/testify/assert"
	"playground/ds/tree"
	"testing"
)

func TestIsSymmetry(t *testing.T) {
	t1 := MinHeightBST([]int{1, 2, 1})
	assert.True(t, IsSymmetry(t1))
	assert.True(t, isEqual(t1, Reverse(Reverse(t1))))

	t2 := MinHeightBST([]int{1, 2, 3, 4, 3, 2, 1})
	assert.True(t, IsSymmetry(t2))
	assert.True(t, isEqual(t2, Reverse(Reverse(t2))))

	t3 := MinHeightBST([]int{1, 2, 3, 4, 3, 2, 1, 4})
	assert.False(t, IsSymmetry(t3))
	assert.True(t, isEqual(t3, Reverse(Reverse(t3))))

	t4 := MinHeightBST([]int{1})
	assert.False(t, IsSymmetry(t4))

	t5 := MinHeightBST([]int{})
	assert.False(t, IsSymmetry(t5))

	t6 := &tree.BTNode{
		Data: 1,
		Left: &tree.BTNode{
			Data: 2,
			Left: &tree.BTNode{
				Data: 3,
			},
		},
		Right: &tree.BTNode{
			Data: 2,
			Right: &tree.BTNode{
				Data: 3,
			},
		},
	}
	assert.True(t, IsSymmetry(t6))
	assert.True(t, isEqual(t6, Reverse(Reverse(t6))))

	t7 := &tree.BTNode{
		Data: 1,
		Left: &tree.BTNode{
			Data: 2,
			Left: &tree.BTNode{
				Data: 3,
			},
		},
		Right: &tree.BTNode{
			Data: 2,
			Right: &tree.BTNode{
				Data: 1,
			},
		},
	}
	assert.False(t, IsSymmetry(t7))
	t71 := Reverse(t7)
	t72 := Reverse(t71)
	assert.True(t, isEqual(t7, t72))

}
