package ch02

import (
	"github.com/stretchr/testify/assert"
	"playground/ds"
	"testing"
)

func TestFindKthToLast(t *testing.T) {
	l1 := ds.LinkedList{}
	l1.AddNode(1) //5//last//first to the last
	l1.AddNode(2) //4//second to the last
	l1.AddNode(3) //3//
	l1.AddNode(4) //2
	l1.AddNode(5) //1//first
	s := FindKthToLast(l1.Head, 1)
	l1.Head = s
	assert.Equal(t, 1, s.Data)
}

func TestFindKthToLastIterative(t *testing.T) {
	l1 := ds.LinkedList{}
	l1.AddNode(1) //1
	l1.AddNode(2) //2
	l1.AddNode(3) //3//
	l1.AddNode(4) //2
	l1.AddNode(5) //1
	s := FindKthToLastIterative(l1.Head, 4)
	assert.Equal(t, 4, s.Data)
}

//func TestFindKthToFirst(t *testing.T) {
//	l1 := ds.LinkedList{}
//	l1.AddNode(1) //1
//	l1.AddNode(2) //2
//	l1.AddNode(3) //3//
//	l1.AddNode(2) //2
//	l1.AddNode(1) //1
//	s := FindKthToFirst(l1.Head, 3)
//	l1.Head = s
//	assert.Equal(t, "[3]->[2]->[1]", l1.String())
//}
