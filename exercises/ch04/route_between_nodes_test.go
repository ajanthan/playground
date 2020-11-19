package ch04

import (
	"github.com/stretchr/testify/assert"
	"playground/ds/graph"
	"testing"
)

func TestHasRoute(t *testing.T) {
	e := &graph.Node{Data: 2}
	s := &graph.Node{
		Data: 5,
		Neighbours: []*graph.Node{
			{
				Data: 1,
				Neighbours: []*graph.Node{
					{Data: 8}, e,
				},
			},
			{
				Data: 2,
				Neighbours: []*graph.Node{
					{Data: 3}, {Data: 2},
				},
			},
		},
	}
	assert.True(t, HasRoute(s, e))
}
