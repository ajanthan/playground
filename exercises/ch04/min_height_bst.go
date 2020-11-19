package ch04

import "playground/ds/graph"

func MinHeightBST(in []int) *graph.Node {
	//1,2
	if len(in) == 1 {
		return &graph.Node{Data: in[0]}
	} else if len(in) == 0 {
		return nil
	}
	root := &graph.Node{}
	mid := len(in) / 2
	root.Data = in[mid]
	left := MinHeightBST(in[:mid])
	right := MinHeightBST(in[mid+1:])
	root.Neighbours = append(root.Neighbours, left, right)
	return root
}
