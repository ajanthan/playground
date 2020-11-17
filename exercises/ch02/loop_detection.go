package ch02

import (
	"playground/ds"
)

func DetectLoop(head *ds.Node) *ds.Node {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow, fast = slow.Next, fast.Next.Next
		if slow == fast {
			break
		}
	}
	slow = head
	for fast != nil && fast.Next != nil {
		if slow == fast {
			return fast
		}
		slow, fast = slow.Next, fast.Next
	}
	return nil
}
