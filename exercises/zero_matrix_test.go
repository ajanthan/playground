package exercises

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestZeroMatrix(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 0, 0}, {3, 4, 0, 0}, {0, 0, 0, 0}, {0, 0, 0, 0}}, ZeroMatrix([][]int{{1, 2, 3, 3}, {3, 4, 1, 3}, {1, 2, 0, 5}, {7, 8, 9, 0}}))
	assert.Equal(t, [][]int{{0, 0, 0, 0}, {0, 4, 1, 0}, {0, 2, 3, 0}, {0, 0, 0, 0}}, ZeroMatrix([][]int{{0, 2, 3, 3}, {3, 4, 1, 3}, {1, 2, 3, 5}, {7, 8, 9, 0}}))
	assert.Equal(t, [][]int{{1, 0, 0, 3}, {3, 0, 0, 3}, {0, 0, 0, 0}, {7, 0, 0, 4}}, ZeroMatrix([][]int{{1, 2, 3, 3}, {3, 4, 1, 3}, {1, 0, 0, 5}, {7, 8, 9, 4}}))
}
