package ds

import "errors"

type Queue []string

func (q *Queue) Add(val string) {
	*q = append(*q, val)
}

func (q *Queue) Remove() (string, error) {
	tq := *q
	if tq.IsEmpty() {
		return "", errors.New("empty queue")
	}
	r := tq[0]
	tq = tq[1:]
	*q = tq
	return r, nil
}

func (q *Queue) Empty() {
	*q = Queue{}
}

func (q *Queue) Peek() (string, error) {
	tq := *q
	if tq.IsEmpty() {
		return "", errors.New("empty queue")
	}
	return tq[0], nil
}
func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
