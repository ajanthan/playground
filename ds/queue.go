package ds

import "errors"

type Queue []string

func (q *Queue) Put(val string) {
	*q = append(*q, val)
}

func (q *Queue) Get() (string, error) {
	tq := *q
	if len(tq) == 0 {
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
