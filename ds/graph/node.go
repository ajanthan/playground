package graph

type Node struct {
	Data       interface{}
	Neighbours []*Node
	Visited    bool
	Marked     bool
}

func DFSearch(n *Node, visit func(n *Node)) {
	if n == nil {
		return
	}
	visit(n)
	n.Visited = true
	for _, neighbour := range n.Neighbours {
		if !neighbour.Visited {
			DFSearch(neighbour, visit)
		}
	}
}

func BFSearch(root *Node, visit func(n *Node)) {
	neighbourQueue := Queue{}
	if root == nil {
		return
	}
	root.Marked = true
	neighbourQueue.Add(root)

	for !neighbourQueue.IsEmpty() {
		node, _ := neighbourQueue.Remove()
		visit(node)
		for _, neighbour := range node.Neighbours {
			if !neighbour.Marked {
				neighbour.Marked = true
				neighbourQueue.Add(neighbour)
			}
		}
	}
}
