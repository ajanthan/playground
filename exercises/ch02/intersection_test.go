package ch02

import (
	"github.com/magiconair/properties/assert"
	"playground/ds"
	"testing"
)

func TestIsIntersection(t *testing.T) {
	l1 := &ds.LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	iNode := l1.Head //3
	l1.AddNode(2)
	l1.AddNode(1)

	l2 := &ds.LinkedList{}
	l2.Head = iNode
	l2.AddNode(1)
	l2.AddNode(2)
	//t.Log("l1", l1.String())
	//t.Log("l2", l2.String())
	assert.Equal(t, iNode, GetIntersectingNode(l1.Head, l2.Head), "Testing equal size list")
	l2.AddNode(3)
	//t.Log("l1", l1.String())
	//t.Log("l2", l2.String())
	assert.Equal(t, iNode, GetIntersectingNode(l1.Head, l2.Head), "Testing unequal size list(l2>l1)")
	l1.AddNode(3)
	l1.AddNode(2)
	//t.Log("l1", l1.String())
	//t.Log("l2", l2.String())
	assert.Equal(t, iNode, GetIntersectingNode(l1.Head, l2.Head), "Testing unequal size list(l1>l2)")
}
