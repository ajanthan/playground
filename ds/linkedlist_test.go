package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := &LinkedList{}
	l.AddNode(4)
	assert.Equal(t, 4, l.Head.Data)
	l.AddNode(3)
	assert.Equal(t, 3, l.Head.Data)
	l.AddNode(2)
	assert.Equal(t, 2, l.Head.Data)
	l.AddNode(1)
	assert.Equal(t, 1, l.Head.Data)
	t.Log(l)
	l.DeleteNode(1)
	assert.Equal(t, 2, l.Head.Data)
	l.DeleteNode(3)
	assert.Equal(t, 2, l.Head.Data)
	l.DeleteNode(4)
	assert.Equal(t, 2, l.Head.Data)
}

func TestSize(t *testing.T) {
	l := &LinkedList{}
	l.AddNode(0)
	l.AddNode(1)
	l.AddNode(2)
	l.AddNode(3)
	assert.Equal(t, 4, Size(l.Head))
	l.DeleteNode(2)
	assert.Equal(t, 3, Size(l.Head))
}

func TestEqual(t *testing.T) {
	l1 := &LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	l1.AddNode(4)
	l2 := &LinkedList{}
	l2.AddNode(1)
	l2.AddNode(2)
	l2.AddNode(3)
	l2.AddNode(4)
	assert.True(t, Equal(l1.Head, l2.Head))
	l2.DeleteNode(3)
	l2.AddNode(3)
	assert.False(t, Equal(l1.Head, l2.Head))
	l1.DeleteNode(1)
	assert.False(t, Equal(l1.Head, l2.Head))
}

func TestCopy(t *testing.T) {
	l1 := &LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	l1.AddNode(4)
	n2 := Clone(l1.Head)
	assert.True(t, Equal(l1.Head, n2))
}

func TestLinkedList_AddAll(t *testing.T) {
	//4,3,2,1
	l1 := &LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	l1.AddNode(4)
	//4,3,2,1
	l2 := &LinkedList{}
	l2.AddNode(1)
	l2.AddNode(2)
	l2.AddNode(3)
	l2.AddNode(4)
	//l1+l2
	//4,3,2,1,4,3,2,1
	l3 := &LinkedList{}
	l3.AddNode(1)
	l3.AddNode(2)
	l3.AddNode(3)
	l3.AddNode(4)
	l3.AddNode(1)
	l3.AddNode(2)
	l3.AddNode(3)
	l3.AddNode(4)
	l1.AddAll(l2)
	assert.True(t, Equal(l1.Head, l3.Head))
}

func TestLinkedList_RemoveFirst(t *testing.T) {
	l1 := &LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	l1.AddNode(4)
	//first=4
	i := l1.RemoveFirst()
	assert.Equal(t, 4, i)
	assert.Equal(t, 3, l1.Head.Data)
}
func TestLinkedList_RemoveLast(t *testing.T) {
	l1 := &LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	l1.AddNode(4)
	//last=1
	i := l1.RemoveLast()
	assert.Equal(t, 1, i)
	i = l1.RemoveLast()
	assert.Equal(t, 2, i)
	l1.AddLast(1)
	i = l1.RemoveLast()
	assert.Equal(t, 1, i)
}
