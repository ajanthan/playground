package ch03

import (
	"playground/ds"
	"strings"
)

func ReverseV0(s string) string {
	stack := &ds.Stack{}
	for _, c := range s {
		stack.Push(string(c))
	}
	var builder strings.Builder

	for !stack.IsEmpty() {
		c, _ := stack.Pop()
		builder.WriteString(c)
	}
	return builder.String()
}
func ReverseV1(s string) string {
	if len(s) <= 1 {
		return s
	}
	var builder strings.Builder
	c := s[0:1]
	r := ReverseV1(s[1:])
	builder.WriteString(r)
	builder.WriteString(c)
	return builder.String()
}
