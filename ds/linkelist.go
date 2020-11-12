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
