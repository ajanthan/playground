package graph

import "errors"

type Queue []*Node

func (q *Queue) Add(val *Node) {
	*q = append(*q, val)
}

func (q *Queue) Remove() (*Node, error) {
	tq := *q
	if tq.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	r := tq[0]
	tq = tq[1:]
	*q = tq
	return r, nil
}

func (q *Queue) Empty() {
	*q = Queue{}
}

func (q *Queue) Peek() (*Node, error) {
	tq := *q
	if tq.IsEmpty() {
		return nil, errors.New("empty queue")
	}
	return tq[0], nil
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
