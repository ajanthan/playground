package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLinkedList(t *testing.T) {
	l := &LinkedList{}
	l.AddNode(4)
	assert.Equal(t, 4, l.head.data)
	l.AddNode(3)
	assert.Equal(t, 3, l.head.data)
	l.AddNode(2)
	assert.Equal(t, 2, l.head.data)
	l.AddNode(1)
	assert.Equal(t, 1, l.head.data)
	t.Log(l)
	l.DeleteNode(1)
	assert.Equal(t, 2, l.head.data)
	l.DeleteNode(3)
	assert.Equal(t, 2, l.head.data)
	l.DeleteNode(4)
	assert.Equal(t, 2, l.head.data)
}
