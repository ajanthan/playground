package ch02

import "playground/ds"

func SumList(l1, l2 ds.LinkedList) *ds.LinkedList {
	var tail *ds.Node
	var head *ds.Node
	n1, n2 := l1.Head, l2.Head
	i, j, c := 0, 0, 0
	for n1 != nil || n2 != nil {
		if n1 == nil {
			i = 0
		} else {
			i = n1.Data
			n1 = n1.Next
		}
		if n2 == nil {
			j = 0
		} else {
			j = n2.Data
			n2 = n2.Next
		}
		node := &ds.Node{}
		node.Data = (i + j + c) % 10
		if head == nil {
			head = node
			tail = node
		} else {
			tail.Next = node
			tail = node
		}
		c = (i + j + c) / 10
	}
	return &ds.LinkedList{
		Head: head,
	}
}
func SumListForward(l1, l2 ds.LinkedList) *ds.LinkedList {
	node, carry := sumListForward(l1.Head, l2.Head)
	if carry > 0 {
		nextNode := &ds.Node{}
		nextNode.Data = carry
		nextNode.Next = node
		node = nextNode
	}
	return &ds.LinkedList{
		Head: node,
	}
}
func sumListForward(l1, l2 *ds.Node) (*ds.Node, int) {
	if l1.Next == nil && l2.Next == nil {
		node := &ds.Node{}
		node.Data = (l1.Data + l2.Data) % 10
		return node, (l1.Data + l2.Data) / 10
	}
	node, carry := sumListForward(l1.Next, l2.Next)
	nextNode := &ds.Node{}
	nextNode.Data = (l1.Data + l2.Data + carry) % 10
	nextNode.Next = node
	return nextNode, (l1.Data + l2.Data + carry) / 10
}
