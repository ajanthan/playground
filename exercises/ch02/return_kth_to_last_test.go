package ch02

import (
	"github.com/stretchr/testify/assert"
	"playground/ds"
	"testing"
)

func TestFindKthToLast(t *testing.T) {
	l1 := ds.LinkedList{}
	//l1.AddNode(1) //1
	//l1.AddNode(2) //2
	l1.AddNode(3) //3//
	l1.AddNode(2) //2
	l1.AddNode(1) //1
	s := FindKthToLast(l1.Head, 2)
	l1.Head = s
	assert.Equal(t, "[2]->[3]", l1.String())
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
