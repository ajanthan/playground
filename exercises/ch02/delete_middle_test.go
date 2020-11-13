package ch02

import (
	"github.com/stretchr/testify/assert"
	"playground/ds"
	"testing"
)

func TestGetMiddle(t *testing.T) {
	l1 := ds.LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	l1.AddNode(4)
	l1.AddNode(5)
	middle, err := GetMiddle(l1)
	assert.NoError(t, err)
	assert.Equal(t, 3, middle.Data)

	l2 := ds.LinkedList{}
	l2.AddNode(1)
	l2.AddNode(2)
	l2.AddNode(3)
	l2.AddNode(4)
	l2.AddNode(5)
	l2.AddNode(6)
	middle, err = GetMiddle(l2)
	assert.NoError(t, err)
	assert.Equal(t, 4, middle.Data)
}

func TestDeleteMiddle(t *testing.T) {
	l1 := ds.LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(3)
	l1.AddNode(4)
	l1.AddNode(5)
	middle, err := GetMiddle(l1)
	assert.NoError(t, err)
	DeleteMiddle(middle)
	assert.Equal(t, "[5]->[4]->[2]->[1]", l1.String())

	l2 := ds.LinkedList{}
	l2.AddNode(1)
	l2.AddNode(2)
	l2.AddNode(3)
	l2.AddNode(4)
	l2.AddNode(5)
	l2.AddNode(6)
	middle, err = GetMiddle(l2)
	assert.NoError(t, err)
	DeleteMiddle(middle)
	assert.Equal(t, "[6]->[5]->[3]->[2]->[1]", l2.String())
}
