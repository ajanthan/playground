package ch05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBitClear(t *testing.T) {
	assert.True(t, IsBitClear(0b11001, 2))
	assert.False(t, IsBitClear(0b110011100, 4))
	assert.True(t, IsBitClear(0b1100100110101, 6))
	assert.False(t, IsBitClear(0b1, 0))
	assert.True(t, IsBitClear(0b0, 0))
}

func TestLongestOnesAfterAFlip(t *testing.T) {
	assert.Equal(t, 2, LongestOnesAfterAFlip(0b1))
	assert.Equal(t, 1, LongestOnesAfterAFlip(0b0))
	assert.Equal(t, 2, LongestOnesAfterAFlip(0b10))
	assert.Equal(t, 4, LongestOnesAfterAFlip(0b11010))
	assert.Equal(t, 8, LongestOnesAfterAFlip(0b11011101111))
	assert.Equal(t, 6, LongestOnesAfterAFlip(0b110110010101011110011))
	assert.Equal(t, 6, LongestOnesAfterAFlip(0b11000111110000111000111))
}
