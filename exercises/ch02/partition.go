package ch02

import (
	"playground/ds"
)

func Partition(l *ds.LinkedList, x int) {
	var p1, p2, joint *ds.Node
	node := l.Head
	for node != nil {
		next := node.Next
		if node.Data < x {
			if p1 == nil {
				joint = node
			}
			node.Next = p1
			p1 = node
		} else {
			node.Next = p2
			p2 = node
		}
		node = next
	}
	joint.Next = p2
	l.Head = p1
}
func PartitionClean(l *ds.LinkedList, x int) {
	p1, p2, node := l.Head, l.Head, l.Head
	for node != nil {
		next := node.Next
		if node.Data < x {
			node.Next = p1
			p1 = node
		} else {
			p2.Next = node
			p2 = node
		}
		node = next
	}
	p2.Next = nil
	l.Head = p1
}
