package ds

import (
	"strconv"
	"strings"
)

type Node struct {
	Next *Node
	Data int
}
type LinkedList struct {
	Head *Node
}

func (l *LinkedList) AddNode(data int) {
	newNode := &Node{Data: data}
	if l.Head == nil {
		l.Head = newNode
	} else {
		newNode.Next = l.Head
		l.Head = newNode
	}
}

func (l *LinkedList) AddLast(data int) {
	tail := getTail(l.Head)
	tail.Next = &Node{Data: data}
}
func (l *LinkedList) RemoveFirst() int {
	r := l.Head.Data
	l.Head = l.Head.Next
	return r
}
func (l *LinkedList) RemoveLast() int {
	return RemoveLast(l.Head).Data
}
func (l *LinkedList) DeleteNode(data int) {
	if l.Head != nil {
		if l.Head.Data == data {
			l.Head = l.Head.Next
			return
		}
		node := l.Head
		for node.Next != nil {
			if node.Next.Data == data {
				node.Next = node.Next.Next
			} else {
				node = node.Next
			}
		}
	}
}
func (l *LinkedList) String() string {
	//var builder strings.Builder
	//builder.WriteString("Head[" + strconv.Itoa(l.Head.Data) + "]")
	//node := l.Head.Next
	//for node != nil {
	//	builder.WriteString("->[" + strconv.Itoa(node.Data) + "]")
	//	node = node.Next
	//}
	//return builder.String()
	return String(l.Head)
}
func String(node *Node) string {
	var builder strings.Builder
	if node.Next == nil {
		builder.WriteString("[" + strconv.Itoa(node.Data) + "]")
		return builder.String()
	} else {
		builder.WriteString("[" + strconv.Itoa(node.Data) + "]->")
		return builder.String() + String(node.Next)
	}
}

func Size(n *Node) int {
	if n == nil {
		return 0
	}
	return 1 + Size(n.Next)
}

func Equal(n1, n2 *Node) bool {
	if Size(n1) != Size(n2) {
		return false
	}
	for n1.Next != nil {
		if n1.Data != n2.Data {
			return false
		}
		n1, n2 = n1.Next, n2.Next
	}
	return true
}
func Clone(n *Node) *Node {
	if n == nil {
		return nil
	}
	node := Clone(n.Next)
	newNode := &Node{}
	newNode.Data = n.Data
	newNode.Next = node
	return newNode

}

func (l *LinkedList) AddAll(another *LinkedList) {
	if l.Head == nil {
		l.Head = another.Head
	} else {
		tail := getTail(l.Head)
		tail.Next = another.Head
	}
}

func getTail(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.Next == nil {
		return n
	}
	return getTail(n.Next)
}

func RemoveLast(n *Node) *Node {
	if n == nil {
		return nil
	}
	if n.Next.Next == nil {
		n.Next = nil
		return n
	}
	return RemoveLast(n.Next)
}
