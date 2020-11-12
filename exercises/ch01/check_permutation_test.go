package ch01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckPermutation(t *testing.T) {
	assert.True(t, CheckPermutation("abcd", "bacd"))
	assert.True(t, CheckPermutation("abcd", "acbd"))
	assert.True(t, CheckPermutation("abcd", "abdc"))
	assert.True(t, CheckPermutation("aacd", "aadc"))

	assert.True(t, CheckPermutation("abcd", "abcd"))

	assert.False(t, CheckPermutation("abcd", "abcde"))
	assert.False(t, CheckPermutation("abcd", "aacd"))
	assert.False(t, CheckPermutation("abcd", "aaaa"))
}
