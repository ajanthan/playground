package ch02

import (
	"github.com/stretchr/testify/assert"
	"playground/ds"

	"testing"
)

func TestDetectLoop(t *testing.T) {
	l1 := &ds.LinkedList{}
	assert.Nil(t, DetectLoop(l1.Head), "nil check")
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	l1.AddNode(4)
	l1.AddNode(5)
	assert.Nil(t, DetectLoop(l1.Head), "negative test")

	l2 := &ds.LinkedList{}
	l2.AddNode(1)
	iNode := l2.Head //1
	l2.AddNode(2)
	l2.AddNode(3)
	iNode.Next = l2.Head
	l2.AddNode(2)
	l2.AddNode(4)
	assert.Equal(t, iNode.Next, DetectLoop(l2.Head), "positive")
}
