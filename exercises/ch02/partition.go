package ch02

import (
	"playground/ds"
)

func Partition(l *ds.LinkedList, x int) {
	var p1 *ds.Node
	var p2 *ds.Node
	var joint *ds.Node
	node := l.Head

	for node != nil {
		if node.Data < x {
			new := &ds.Node{}
			if p1 == nil {
				joint = new
			}
			new.Data = node.Data
			new.Next = p1
			p1 = new
		} else {
			new := &ds.Node{}
			new.Data = node.Data
			new.Next = p2
			p2 = new
		}
		node = node.Next
	}
	joint.Next = p2
	l.Head = p1
}
