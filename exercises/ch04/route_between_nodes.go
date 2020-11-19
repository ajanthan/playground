package ch04

import "playground/ds/graph"

func HasRoute(s *graph.Node, e *graph.Node) bool {
	queue := &graph.Queue{}
	if s == nil {
		return false
	}
	s.Marked = true
	queue.Add(s)

	for !queue.IsEmpty() {
		node, _ := queue.Remove()
		if node == e {
			return true
		}
		for _, neighbour := range node.Neighbours {
			if !neighbour.Marked {
				neighbour.Marked = true
				queue.Add(neighbour)
			}
		}
	}
	return false
}
