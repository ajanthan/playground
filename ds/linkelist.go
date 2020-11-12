package ds

import (
	"strconv"
	"strings"
)

type Node struct {
	next *Node
	data int
}
type LinkedList struct {
	head *Node
}

func (l *LinkedList) AddNode(data int) {
	newNode := &Node{data: data}
	if l.head == nil {
		l.head = newNode
	} else {
		newNode.next = l.head
		l.head = newNode
	}
}
func (l *LinkedList) DeleteNode(data int) {
	if l.head != nil {
		if l.head.data == data {
			l.head = l.head.next
			return
		}
		node := l.head
		for node.next != nil {
			if node.next.data == data {
				node.next = node.next.next
			} else {
				node = node.next
			}
		}
	}
}
func (l *LinkedList) String() string {
	//var builder strings.Builder
	//builder.WriteString("head[" + strconv.Itoa(l.head.data) + "]")
	//node := l.head.next
	//for node != nil {
	//	builder.WriteString("->[" + strconv.Itoa(node.data) + "]")
	//	node = node.next
	//}
	//return builder.String()
	return String(l.head)
}
func String(node *Node) string {
	var builder strings.Builder
	if node.next == nil {
		builder.WriteString("[" + strconv.Itoa(node.data) + "]")
		return builder.String()
	} else {
		builder.WriteString("[" + strconv.Itoa(node.data) + "]->")
		return builder.String() + String(node.next)
	}
}
