package ch01

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsPalindromePermutation(t *testing.T) {
	assert.True(t, IsPalindromePermutation("Tact Coa"))
	assert.True(t, IsPalindromePermutation("oCct aTa"))
	assert.True(t, IsPalindromePermutation("zooz"))
	assert.True(t, IsPalindromePermutation("Tacta Coaa"))

	assert.False(t, IsPalindromePermutation("Tacta Coa"))
	assert.False(t, IsPalindromePermutation("Tact Coat"))
}
