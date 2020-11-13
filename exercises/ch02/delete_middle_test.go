package ch02

import (
	"github.com/stretchr/testify/assert"
	"playground/ds"
	"testing"
)

func TestDeleteMiddle(t *testing.T) {
	l1 := ds.LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	l1.AddNode(4)
	l1.AddNode(5)
	DeleteMiddle(l1)
	assert.Equal(t, "[5]->[4]->[2]->[1]", l1.String())

	l2 := ds.LinkedList{}
	l2.AddNode(1)
	l2.AddNode(2)
	l2.AddNode(3)
	l2.AddNode(4)
	l2.AddNode(5)
	l2.AddNode(6)
	DeleteMiddle(l2)
	assert.Equal(t, "[6]->[5]->[3]->[2]->[1]", l2.String())
}
