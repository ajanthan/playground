package ch02

import (
	"playground/ds"
)

func Reverse(l ds.LinkedList) ds.LinkedList {
	r := ds.LinkedList{}
	cur := l.Head.Next
	l.Head.Next = nil
	head := reverse(l.Head, cur)
	r.Head = head
	return r
}

func reverse(prev, cur *ds.Node) *ds.Node {
	if cur == nil {
		return prev
	}
	next := cur.Next
	cur.Next = prev
	return reverse(cur, next)

}
