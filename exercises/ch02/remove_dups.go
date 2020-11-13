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
	cNode := l.Head //n[1]->[2]->[3]->[2]->[1]
	//	pNode := l.Head
	for cNode.Next != nil {
		_, ok := ht[cNode.Next.Data]
		if ok {
			//pNode.Next = cNode.Next
			cNode.Next = cNode.Next.Next
		} else {
			ht[cNode.Next.Data] = cNode.Next.Data
			//pNode = cNode
			cNode = cNode.Next
		}

	}
	return nil
}

func RemoveDupsWith2Pointer(l ds.LinkedList) error {
	if l.Head == nil {
		return errors.New("Empty")
	}
	cNode := l.Head
	for cNode != nil {
		runner := cNode
		for runner.Next != nil {
			if runner.Next.Data == cNode.Data {
				runner.Next = runner.Next.Next
			} else {
				runner = runner.Next
			}
		}
		cNode = cNode.Next
	}
	return nil
}
