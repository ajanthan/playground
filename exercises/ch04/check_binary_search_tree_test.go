package ch04

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"playground/ds/tree"
	"testing"
)

func TestIsBST(t *testing.T) {
	cases := []struct {
		d string
		f func(root *tree.BTNode) bool
	}{{
		d: "IsBSTV1",
		f: IsBSTV1,
	},
		{
			d: "IsBSTV12",
			f: IsBSTV2,
		},
	}
	for i, c := range cases {
		in1 := []int{1, 2, 3, 4, 5, 6, 7}
		root := MinHeightBST(in1)
		assert.True(t, c.f(root), fmt.Sprintf("Case#%d:%s", i, c.d))
		root.Left.Right.Data = 5
		assert.False(t, c.f(root), fmt.Sprintf("Case#%d:%s", i, c.d))

		in2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
		root = MinHeightBST(in2)
		assert.True(t, c.f(root), fmt.Sprintf("Case#%d:%s", i, c.d))
		root.Right.Right.Left.Data = 7
		assert.False(t, c.f(root), fmt.Sprintf("Case#%d:%s", i, c.d))
	}
}
func TestBSTtoSlice(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5, 6, 7}
	root := MinHeightBST(in1)
	assert.Equal(t, in1, BSTtoSlice(root))

	in2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}
	root = MinHeightBST(in2)
	assert.Equal(t, in2, BSTtoSlice(root))
}
