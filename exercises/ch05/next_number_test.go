package ch05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNextNumber(t *testing.T) {
	s, l := NextNumber(0b1)
	assert.Equal(t, 0b1, s)
	assert.Equal(t, 0b1, l)

	s, l = NextNumber(0b0)
	assert.Equal(t, 0b0, s)
	assert.Equal(t, 0b0, l)

	s, l = NextNumber(0b1010)
	assert.Equal(t, 0b0011, s)
	assert.Equal(t, 0b1100, l)

	s, l = NextNumber(0b10100011011)
	assert.Equal(t, 0b00000111111, s)
	assert.Equal(t, 0b11111100000, l)
}
