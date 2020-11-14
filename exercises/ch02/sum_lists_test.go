package ch02

import (
	"github.com/magiconair/properties/assert"
	"playground/ds"
	"testing"
)

func TestSumList(t *testing.T) {
	l1 := ds.LinkedList{}
	l1.AddNode(6)
	l1.AddNode(1)
	l1.AddNode(7)
	t.Log(l1.String())
	l2 := ds.LinkedList{}
	l2.AddNode(2)
	l2.AddNode(9)
	l2.AddNode(5)
	t.Log(l2.String())
	l3 := SumList(l1, l2)
	assert.Equal(t, l3.String(), "[2]->[1]->[9]")

	l4 := ds.LinkedList{}
	l4.AddNode(6)
	l4.AddNode(1)
	l4.AddNode(7)
	l4.AddNode(8)
	t.Log(l4.String())
	l5 := ds.LinkedList{}
	l5.AddNode(2)
	l5.AddNode(9)
	l5.AddNode(5)
	t.Log(l5.String())
	l6 := SumList(l4, l5)
	t.Log(l6.String())
	assert.Equal(t, l6.String(), "[3]->[7]->[4]->[6]")
}
