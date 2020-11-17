package ch03

import "errors"

type IntStack struct {
	slice []int
	min   []int
}

func (i *IntStack) Push(in int) {
	if len(i.slice) == 0 {
		i.min = append(i.min, in)
	} else {
		min := i.min[len(i.min)-1]
		if min >= in {
			i.min = append(i.min, in)
		}
	}
	i.slice = append(i.slice, in)
}

func (i *IntStack) Pop() (int, error) {
	if len(i.slice) == 0 {
		return 0, errors.New("empty")
	}
	r := i.slice[len(i.slice)-1]
	min, err := i.Min()
	if err != nil {
		return 0, err
	}
	if min == r {
		ms := make([]int, 0)
		i.min = append(ms, i.min[:len(i.min)-1]...)
	}
	ns := make([]int, 0)
	i.slice = append(ns, i.slice[:len(i.slice)-1]...)
	return r, nil
}
func (i *IntStack) Min() (int, error) {
	if len(i.slice) == 0 {
		return 0, errors.New("empty")
	}
	return i.min[len(i.min)-1], nil
}
