package ch02

import (
	"github.com/stretchr/testify/assert"
	"playground/ds"
	"testing"
)

func TestPartition(t *testing.T) {
	l1 := ds.LinkedList{}
	l1.AddNode(1)
	l1.AddNode(2)
	l1.AddNode(10)
	l1.AddNode(5)
	l1.AddNode(8)
	l1.AddNode(5)
	l1.AddNode(3)
	Partition(&l1, 5)
	assert.Equal(t, "[1]->[2]->[3]->[10]->[5]->[8]->[5]", l1.String())
}
