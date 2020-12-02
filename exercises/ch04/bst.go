package ch04

import (
	"math"
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
	insert(b.Root, i)
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
			root.Left = &tree.BTNode{Data: i}
			return
		}
		insert(root.Left, i)
	}
	if i > root.Data.(int) {
		if root.Right == nil {
			root.Right = &tree.BTNode{Data: i}
			return
		}
		insert(root.Right, i)
	}
}

func (b *BST) GetRandomNode() int {
	seed := rand.NewSource(time.Now().UnixNano())
	random := rand.New(seed)
	r, _ := getKthNode(b.Root, random.Intn(b.Len))
	return r
}
func getKthNode(root *tree.BTNode, k int) (int, int) {
	if root == nil {
		return math.MinInt64, k
	}
	if k == 0 {
		return root.Data.(int), k
	}
	k--
	left, k := getKthNode(root.Left, k)
	if left != math.MinInt64 {
		return left, k
	}
	right, k := getKthNode(root.Right, k)
	if right != math.MinInt64 {
		return right, k
	}
	return math.MinInt64, k
}
