package ch03

import "errors"

type IntStack struct {
	slice []int
	size  int
}
type SubStack struct {
	stack []IntStack
}

func (i *IntStack) Push(in int) error {
	if len(i.slice) > i.size {
		return errors.New("stackoverflow")
	}
	i.slice = append(i.slice, in)
	return nil
}
func (i *IntStack) IsEmpty() bool {
	return len(i.slice) == 0
}

func (i *IntStack) Pop() (int, error) {
	if len(i.slice) == 0 {
		return 0, errors.New("empty")
	}
	r := i.slice[len(i.slice)-1]
	ns := make([]int, 0)
	i.slice = append(ns, i.slice[:len(i.slice)-1]...)
	return r, nil
}

func (s *SubStack) Push(in int) error {
	if len(s.stack) == 0 {
		s.stack = append(s.stack, IntStack{size: 10})
	}
	err := s.stack[len(s.stack)-1].Push(in)
	if err != nil {
		s.stack = append(s.stack, IntStack{size: 10})
		err = s.stack[len(s.stack)-1].Push(in)
		return err
	}
	return nil
}

func (s *SubStack) Pop() (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("empty")
	}
	r, err := s.stack[len(s.stack)-1].Pop()
	if s.stack[len(s.stack)-1].IsEmpty() {
		ss := make([]IntStack, 0)
		s.stack = append(ss, s.stack[:len(s.stack)-1]...)
	}
	return r, err
}

func (s *SubStack) PopAt(index int) (int, error) {
	if len(s.stack) == 0 {
		return 0, errors.New("empty")
	}
	if index >= len(s.stack) {
		return 0, errors.New("index out of bound")
	}
	r, err := s.stack[index].Pop()
	if s.stack[index].IsEmpty() {
		ss := append(s.stack[:index], s.stack[index+1:]...)
		s.stack = ss
	}
	return r, err
}
