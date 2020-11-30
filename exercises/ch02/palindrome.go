package ch02

import (
	"playground/ds"
)

func IsPalindromeV0(l ds.LinkedList) bool {
	if l.Head == nil {
		return false
	}
	n := ds.Clone(l.Head)
	l2 := Reverse(ds.LinkedList{Head: n})
	return ds.Equal(l2.Head, l.Head)
}
func IsPalindromeV1(l ds.LinkedList) bool {
	var stack []int
	slow := l.Head
	fast := l.Head
	i := -1

	//foundMid := false
	//for slow != nil {
	//	foundMid = fast == nil || fast.Next == nil
	//	if foundMid {
	//		if stack[i] != slow.Data {
	//			return false
	//		}
	//		i--
	//	} else {
	//		stack = append(stack, slow.Data)
	//		i++
	//		fast = fast.Next.Next
	//		if fast != nil && fast.Next == nil {
	//			slow = slow.Next
	//		}
	//	}
	//	slow = slow.Next
	//}
	for fast != nil && fast.Next != nil {
		stack = append(stack, slow.Data)
		i++
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil && fast.Next == nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Data != stack[i] {
			return false
		}
		i--
		slow = slow.Next
	}
	return true
}
func IsPalindromeV2(l ds.LinkedList) bool {
	_, isPal := IsPalindrome(l.Head, ds.Size(l.Head))
	return isPal
}
func IsPalindrome(n *ds.Node, s int) (*ds.Node, bool) {
	if s == 0 {
		return n, true
	} else if s == 1 {
		return n.Next, true
	}
	node, isPal := IsPalindrome(n.Next, s-2)
	if !isPal {
		return node.Next, isPal
	}
	if node.Data == n.Data {
		return node.Next, true
	}
	return node.Next, false
}
