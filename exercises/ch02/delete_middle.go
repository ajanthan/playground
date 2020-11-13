package ch02

import (
	"errors"
	"playground/ds"
)

func GetMiddle(l ds.LinkedList) (*ds.Node, error) {
	//l=[1]->[2]->[3]
	count := 0
	p1 := l.Head
	p2 := l.Head
	for p1 != nil {
		count++      //1   2   3
		p1 = p1.Next //[2] [3]  nil
	}
	if count < 3 {
		return nil, errors.New("invalid input")
	}
	middle := 0
	if count%2 == 0 {
		middle = count / 2
	} else {
		middle = count/2 + 1 //middle=2
	}
	for i := 2; i <= count; i++ { //count=3
		if i == middle {
			break
		}
		p2 = p2.Next
	}
	return p2.Next, nil
}

func DeleteMiddle(node *ds.Node) error {
	if node.Next == nil {
		return errors.New("not a middle node")
	}
	node.Data = node.Next.Data
	node.Next = node.Next.Next
	return nil
}
