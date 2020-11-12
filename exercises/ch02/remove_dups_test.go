package ch02

import (
	"github.com/stretchr/testify/assert"
	"playground/ds"
	"testing"
)

func TestRemoveDups(t *testing.T) {
	l1 := ds.LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(2)
	l1.AddNode(1)
	l1.AddNode(5)
	RemoveDups(l1)
	assert.Equal(t, "[5]->[1]->[2]", l1.String())

	l2 := ds.LinkedList{}
	l2.AddNode(5)
	l2.AddNode(1)
	l2.AddNode(2)
	l2.AddNode(2)
	l2.AddNode(1)
	l2.AddNode(5)
	RemoveDups(l2)
	assert.Equal(t, "[5]->[1]->[2]", l2.String())

	l3 := ds.LinkedList{}
	l3.AddNode(1)
	l3.AddNode(1)
	l3.AddNode(1)
	l3.AddNode(1)
	l3.AddNode(1)
	l3.AddNode(1)
	RemoveDups(l3)
	assert.Equal(t, "[1]", l3.String())
}
