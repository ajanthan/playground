package ch02

import (
	"github.com/stretchr/testify/assert"
	"playground/ds"
	"testing"
)

func TestReverse(t *testing.T) {
	l1 := ds.LinkedList{}
	l1.AddNode(1) //1
	l1.AddNode(2) //2
	l1.AddNode(3) //3//
	l1.AddNode(4) //2
	s := Reverse(l1)
	assert.Equal(t, "[1]->[2]->[3]->[4]", s.String())
}
