package graph

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestNode_String(t *testing.T) {
	root := &Node{
		Data: 5,
		Neighbours: []*Node{
			{
				Data: 1,
				Neighbours: []*Node{
					{Data: 8}, {Data: 2},
				},
			},
			{
				Data: 2,
				Neighbours: []*Node{
					{Data: 3}, {Data: 2},
				},
			},
		},
	}
	min := root.Data.(int)
	findMin := func(n *Node) {
		if n.Data.(int) < min {
			min = n.Data.(int)
		}
	}
	DFSearch(root, findMin)
	assert.Equal(t, 1, min)

	max := root.Data.(int)
	findMax := func(n *Node) {
		if n.Data.(int) > max {
			max = n.Data.(int)
		}
	}
	BFSearch(root, findMax)
	assert.Equal(t, 8, max)

}
