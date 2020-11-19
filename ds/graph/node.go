package graph

type Node struct {
	Data       interface{}
	Neighbours []*Node
	visited    bool
	marked     bool
}

func DFSearch(n *Node, visit func(n *Node)) {
	if n == nil {
		return
	}
	visit(n)
	n.visited = true
	for _, neighbour := range n.Neighbours {
		if !neighbour.visited {
			DFSearch(neighbour, visit)
		}
	}
}

func BFSearch(root *Node, visit func(n *Node)) {
	neighbourQueue := Queue{}
	if root == nil {
		return
	}
	root.marked = true
	neighbourQueue.Add(root)

	for !neighbourQueue.IsEmpty() {
		node, _ := neighbourQueue.Remove()
		visit(node)
		for _, neighbour := range node.Neighbours {
			if !neighbour.marked {
				neighbour.marked = true
				neighbourQueue.Add(neighbour)
			}
		}
	}
}
