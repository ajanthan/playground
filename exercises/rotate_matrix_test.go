package exercises

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRotateImage90(t *testing.T) {
	runTestRotateImage90([][]int{{1, 2}, {3, 4}}, t)
	runTestRotateImage90([][]int{{1, 2, 3}, {3, 4, 5}, {9, 7, 5}}, t)
	runTestRotateImage90([][]int{{1, 2, 2, 3}, {3, 4, 4, 2}, {1, 2, 3, 4}, {9, 7, 6, 5}}, t)
}
func runTestRotateImage90(mn [][]int, t *testing.T) {
	assert.Equal(t, mn, RotateImage90(RotateImage90(RotateImage90(RotateImage90(mn)))))
}
