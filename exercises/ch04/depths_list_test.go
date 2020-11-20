package ch04

import (
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestListOfDepths(t *testing.T) {
	in3 := []int{1, 2, 3, 4, 5, 6, 7}
	root := MinHeightBST(in3)
	ret := ListOfDepths(root)
	assert.Equal(t, 4, ret[0][0].Data)
	assert.Equal(t, 2, ret[1][0].Data)
	assert.Equal(t, 6, ret[1][1].Data)
	assert.Equal(t, 1, ret[2][0].Data)
	assert.Equal(t, 3, ret[2][1].Data)
	assert.Equal(t, 5, ret[2][2].Data)
	assert.Equal(t, 7, ret[2][3].Data)
}
