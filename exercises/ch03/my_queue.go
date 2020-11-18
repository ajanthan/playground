package ch03

import "playground/ds"

type MyQueue struct {
	forward ds.Stack
	reverse ds.Stack
}

func NewMyQueue() *MyQueue {
	return &MyQueue{
		forward: ds.Stack{},
		reverse: ds.Stack{},
	}
}

func (m *MyQueue) Add(s string) {
	m.forward.Push(s)
}
func (m *MyQueue) Remove() (string, error) {
	if m.reverse.IsEmpty() {
		for !m.forward.IsEmpty() {
			s, _ := m.forward.Pop()
			m.reverse.Push(s)
		}
	}
	r, err := m.reverse.Pop()
	return r, err
}
func (m *MyQueue) Peek() (string, error) {
	if m.reverse.IsEmpty() {
		for !m.forward.IsEmpty() {
			s, _ := m.forward.Pop()
			m.reverse.Push(s)
		}
	}
	r, err := m.reverse.Peek()
	return r, err
}
func (m *MyQueue) IsEmpty() bool {
	return m.forward.IsEmpty()
}
