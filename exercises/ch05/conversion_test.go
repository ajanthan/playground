package ch05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberOfBitFlipToConvert(t *testing.T) {
	assert.Equal(t, 1, NumberOfBitFlipToConvert(0b0, 0b1))
	assert.Equal(t, 1, NumberOfBitFlipToConvert(0b1, 0b0))
	assert.Equal(t, 0, NumberOfBitFlipToConvert(0b0, 0b0))
	assert.Equal(t, 2, NumberOfBitFlipToConvert(0b11101, 0b01111))
	assert.Equal(t, 5, NumberOfBitFlipToConvert(0b11010101, 0b011))
}
