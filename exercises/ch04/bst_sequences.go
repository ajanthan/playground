package ch04

import (
	"playground/ds"
	"playground/ds/tree"
)

type Result struct {
	Data []*ds.LinkedList
}

func AllSequences(root *tree.BTNode) *Result {
	result := &Result{}
	if root == nil {
		return result
	}
	prefix := &ds.LinkedList{}
	prefix.AddNode(root.Data.(int))

	leftSeq := AllSequences(root.Left)
	rightSeq := AllSequences(root.Right)
	for _, left := range leftSeq.Data {
		for _, right := range rightSeq.Data {
			weaved := &Result{}
			Merge(weaved, prefix, left, right)
			result.Data = append(result.Data, weaved.Data...)
		}
	}
	if len(leftSeq.Data) == 0 && len(rightSeq.Data) == 0 {
		r := &ds.LinkedList{}
		r.AddAll(prefix)
		result.Data = append(result.Data, r)
	}
	return result
}

func Merge(results *Result, prefix *ds.LinkedList, left *ds.LinkedList, right *ds.LinkedList) {
	if ds.Size(left.Head) == 0 || ds.Size(right.Head) == 0 {
		result := &ds.LinkedList{}
		clonedPrefix := &ds.LinkedList{}
		clonedPrefix.Head = ds.Clone(prefix.Head)
		result.AddAll(clonedPrefix)
		result.AddAll(left)
		result.AddAll(right)
		results.Data = append(results.Data, result)
		return
	}
	lHeadFirst := left.RemoveFirst()
	prefix.AddLast(lHeadFirst)
	Merge(results, prefix, left, right)
	prefix.RemoveLast()
	left.AddNode(lHeadFirst)

	rHeadFirst := right.RemoveFirst()
	prefix.AddLast(rHeadFirst)
	Merge(results, prefix, left, right)
	prefix.RemoveLast()
	right.AddNode(rHeadFirst)
}
