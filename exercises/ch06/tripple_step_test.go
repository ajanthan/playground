package ch06

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFindSteps(t *testing.T) {
	w := FindSteps(1)
	assert.Equal(t, 1, w)

	w = FindSteps(2)
	assert.Equal(t, 2, w)

	w = FindSteps(3)
	assert.Equal(t, 3, w)

	w = FindSteps(4)
	assert.Equal(t, 4, w)

	w = FindSteps(5)
	assert.Equal(t, 4, w)

	w = FindSteps(10)
	assert.Equal(t, 14, w)
}
