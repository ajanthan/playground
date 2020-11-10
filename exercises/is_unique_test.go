package exercises

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsUnique(t *testing.T) {

	assert.True(t, IsUnique(""))
	assert.True(t, IsUnique(" "))
	assert.True(t, IsUnique("a"))
	assert.True(t, IsUnique("trieyu"))
	assert.True(t, IsUnique("4l1h3b"))
	assert.True(t, IsUnique("fg?056*&("))

	assert.False(t, IsUnique("aa"))
	assert.False(t, IsUnique("a12344rtt"))
	assert.False(t, IsUnique("*dgkoanmlw*"))
	assert.False(t, IsUnique("*dgkoanmlw1234567890-=qwertyuioasdgjlxcvnm"))
}
