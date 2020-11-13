package ch02

import (
	"playground/ds"
)

//9.46

func FindKthToLast(l *ds.Node, k int) *ds.Node {
	//1->2->3 k=2
	index := 0
	return findKthToLast(l, k, &index)

}
func findKthToLast(l *ds.Node, k int, i *int) *ds.Node {
	if l == nil {
		return nil
	}
	node := findKthToLast(l.Next, k, i)
	*i++
	if *i == k {
		return l
	}
	return node

}

//func FindKthToFirst(l *ds.Node, k int) *ds.Node {
//	//1->2->3 k=2
//	index := 0
//	return findKthToFirst(l, k, &index)
//
//}
//func findKthToFirst(l *ds.Node, k int, i *int) *ds.Node {
//	if l == nil {
//		return nil
//	}
//	*i++
//	fmt.Println("f1 node=", l, " i=", *i, " k=", k)
//	if *i == k {
//		return l
//	} else {
//		l.Next = l
//	}
//	node := findKthToFirst(l.Next, k, i)
//	*i--
//	fmt.Println("f2 node=", node, " i=", *i, " k=", k)
//	if *i == 1 {
//		l.Next = nil
//	}
//	node.Next = l
//	return node
//
//}
