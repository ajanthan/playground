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
