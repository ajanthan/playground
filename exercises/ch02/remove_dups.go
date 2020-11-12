package ch02

import (
	"errors"
	"playground/ds"
)

func RemoveDups(l ds.LinkedList) error {
	//h[1]->[2]->[3]->[2]->[1]
	ht := make(map[int]int)
	if l.Head == nil {
		return errors.New("Empty")
	}
	ht[l.Head.Data] = l.Head.Data
	cNode := l.Head.Next //n[1]->[2]->[3]->[2]->[1]
	pNode := l.Head
	for cNode != nil {
		_, ok := ht[cNode.Data]
		if ok {
			pNode.Next = cNode.Next
		} else {
			ht[cNode.Data] = cNode.Data
			pNode = cNode
		}
		cNode = cNode.Next
	}
	return nil
}
