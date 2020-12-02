package ch04

import (
	"math/rand"
	"playground/ds/tree"
	"time"
)

type BST struct {
	Root *tree.BTNode
	Len  int
}

func (b *BST) Insert(i int) {
	b.Len++
	if b.Root == nil {
		b.Root = &tree.BTNode{Data: i, Size: 1}
	} else {
		insert(b.Root, i)
	}
}
func (b *BST) Find(i int) *tree.BTNode {
	return find(b.Root, i)
}

func find(root *tree.BTNode, i int) *tree.BTNode {
	if root == nil {
		return nil
	}
	if root.Data.(int) == i {
		return root
	} else if root.Data.(int) > i {
		return find(root.Left, i)
	} else {
		return find(root.Right, i)
	}
}

func delete(root *tree.BTNode, i int) *tree.BTNode {
	if root == nil {
		return nil
	}
	if root.Data.(int) == i {
		return root
	} else if root.Data.(int) > i {
		r := delete(root.Left, i)
		if r != nil {
			root.Left = r.Left
			r.Left.Right = r.Right
		}
		return nil
	} else {
		r := delete(root.Right, i)
		if r != nil {
			root.Left = r.Left
			r.Left.Right = r.Right
		}
		return nil
	}
}

func insert(root *tree.BTNode, i int) {
	if root == nil {
		return
	}
	if i <= root.Data.(int) {
		if root.Left == nil {
			root.Left = &tree.BTNode{Data: i, Size: 1}
			root.Size++
			return
		}
		insert(root.Left, i)
		root.Size++
	}
	if i > root.Data.(int) {
		if root.Right == nil {
			root.Right = &tree.BTNode{Data: i, Size: 1}
			root.Size++
			return
		}
		insert(root.Right, i)
		root.Size++
	}
}

func (b *BST) GetRandomNode() int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	r := GetKthNode(b.Root, random.Intn(b.Root.Size))
	return r.Data.(int)
}
func GetKthNode(root *tree.BTNode, k int) *tree.BTNode {
	if root == nil {
		return nil
	}
	if k == 0 {
		return root
	} else {
		if root.Left != nil {
			if k == root.Left.Size {
				return root
			} else if k < root.Left.Size {
				return GetKthNode(root.Left, k-1)
			} else {
				return GetKthNode(root.Right, k-1-root.Left.Size)
			}
		} else {
			return GetKthNode(root.Right, k-1)
		}
	}
}
