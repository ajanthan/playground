package ch04

import (
	"github.com/stretchr/testify/assert"
	"playground/ds/tree"
	"strconv"
	"strings"
	"testing"
)

func TestFindPathsWithWight(t *testing.T) {
	in1 := []int{1, 2, 3, 4, 5, 6, 7}
	root := MinHeightBST(in1)
	/*

	 */
	bt := NewBT(3, 11)
	bt.PrintBT(root, 11/2, 0)
	t.Log(bt)
	paths := FindPathsWithWight(root, 3)
	assert.Equal(t, 2, paths)
}

type BT struct {
	Rows []string
}

func NewBT(r, c int) *BT {
	var builder strings.Builder
	bt := &BT{}
	for i := 0; i < r; i++ {
		builder.Reset()
		for j := 0; j < c; j++ {
			builder.WriteString(" ")
		}
		bt.Rows = append(bt.Rows, builder.String())
	}
	return bt
}
func (bt *BT) PrintBT(root *tree.BTNode, pos, level int) {
	if root == nil {
		return
	}
	row := bt.Rows[level]
	row = row[:pos] + strconv.Itoa(root.Data.(int)) + row[pos-1:]
	bt.Rows[level] = row

	bt.PrintBT(root.Left, pos-(len(bt.Rows)-level-1), level+1)
	bt.PrintBT(root.Right, pos+(len(bt.Rows)+level+1), level+1)
}
func (bt *BT) String() string {
	var builder strings.Builder
	builder.WriteString("\n")
	for _, row := range bt.Rows {
		builder.WriteString(row)
		builder.WriteString("\n")
	}
	return builder.String()
}
