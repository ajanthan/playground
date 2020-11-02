package ds

import "errors"

type Stack struct {
	slice []string
}

func NewStack() *Stack {
	slice := make([]string, 0)
	stack := &Stack{
		slice: slice,
	}
	return stack
}

func (s *Stack) Push(val string) {
	s.slice = append(s.slice, val)
}

func (s *Stack) Pop() (string, error) {
	if len(s.slice) == 0 {
		return "", errors.New("empty stack")
	}
	r := s.slice[len(s.slice)-1]
	s.slice = s.slice[0 : len(s.slice)-1]
	return r, nil
}

func (s *Stack) Clear() {
	s.slice = make([]string, 0)
}
