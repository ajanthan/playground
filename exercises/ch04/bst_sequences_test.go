package ch04

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAllSequences(t *testing.T) {
	in1 := []int{1, 2, 3}
	root := MinHeightBST(in1)
	assert.Equal(t, "[[2]->[1]->[3] [2]->[3]->[1]]", fmt.Sprintf("%s", AllSequences(root).Data))
}
