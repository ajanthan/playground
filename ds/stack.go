package ds

import "errors"

type Stack []string

func (s *Stack) Push(val string) {
	*s = append(*s, val)
}

func (s *Stack) Pop() (string, error) {
	p := *s
	if len(p) == 0 {
		return "", errors.New("empty stack")
	}
	r := p[len(p)-1]
	p = p[0 : len(p)-1]
	*s = p
	return r, nil
}

func (s *Stack) Clear() {
	*s = Stack{}
}
