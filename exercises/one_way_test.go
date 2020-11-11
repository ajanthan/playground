package exercises

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsOneAway(t *testing.T) {
	assert.True(t, IsOneAway("pale", "pale"))
	assert.True(t, IsOneAway("pale", "ple"))
	assert.True(t, IsOneAway("pales", "pale"))
	assert.True(t, IsOneAway("pale", "bale"))
	assert.False(t, IsOneAway("pale", "bake"))
}
