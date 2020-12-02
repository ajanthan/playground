package ch04

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBST_GetRandomNode(t *testing.T) {
	in1 := []int{4, 2, 6, 1, 3, 6, 5, 7}
	bst := &BST{}
	for _, i := range in1 {
		bst.Insert(i)
	}
	i := bst.GetRandomNode()
	assert.Contains(t, in1, i)
}
func Test_GetKthNode(t *testing.T) {
	in1 := []int{4, 2, 6, 1, 3, 6, 5, 7}
	bst := &BST{}
	for _, i := range in1 {
		bst.Insert(i)
	}
	i := GetKthNode(bst.Root, 0).Data.(int)
	assert.Contains(t, in1, i)
	i = GetKthNode(bst.Root, bst.Root.Size-1).Data.(int)
	assert.Contains(t, in1, i)
}

func TestBST_Insert(t *testing.T) {
	in1 := []int{4, 2, 6, 1, 3, 6, 5, 7}
	bst := &BST{}
	for _, i := range in1 {
		bst.Insert(i)
	}
	assert.Equal(t, 8, bst.Root.Size)
}

func TestBST_Find(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5, 6, 7}
	root := MinHeightBST(in1)
	bst := &BST{Root: root, Len: 7}
	assert.Equal(t, 7, bst.Find(7).Data.(int))
	assert.Equal(t, 1, bst.Find(1).Data.(int))
	assert.Equal(t, 4, bst.Find(4).Data.(int))
	assert.Nil(t, bst.Find(9))
}
