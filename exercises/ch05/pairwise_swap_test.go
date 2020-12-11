package ch05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSwapOddEvenBits(t *testing.T) {
	assert.Equal(t, 0b01, SwapOddEvenBits(0b10))
	assert.Equal(t, 0b1000, SwapOddEvenBits(0b0100))
	assert.Equal(t, 0b100001, SwapOddEvenBits(0b010010))
	assert.Equal(t, 0b101001, SwapOddEvenBits(0b010110))
}
