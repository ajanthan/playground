package ch02

import (
	"errors"
	"playground/ds"
)

func DeleteMiddle(l ds.LinkedList) error {
	//l=[1]->[2]->[3]
	count := 0
	p1 := l.Head
	p2 := l.Head
	for p1 != nil {
		count++      //1   2   3
		p1 = p1.Next //[2] [3]  nil
	}
	if count < 3 {
		return errors.New("invalid input")
	}
	middle := 0
	if count%2 == 0 {
		middle = count / 2
	} else {
		middle = count/2 + 1 //middle=2
	}
	for i := 2; i <= count; i++ { //count=3
		if i == middle {
			p2.Next = p2.Next.Next
			return nil
		}
		p2 = p2.Next
	}
	return nil
}
