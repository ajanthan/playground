package ch02

import "playground/ds"

func GetIntersectingNode(n1, n2 *ds.Node) *ds.Node {
	l1 := ds.Size(n1)
	l2 := ds.Size(n2)
	if l1 > l2 {
		for i := 0; i < l1-l2; i++ {
			n1 = n1.Next
		}
	} else if l1 < l2 {
		for i := 0; i < l2-l1; i++ {
			n2 = n2.Next
		}
	}
	for n1 != nil && n2 != nil {
		if n1 == n2 {
			return n1
		}
		n1, n2 = n1.Next, n2.Next
	}
	return nil
}
