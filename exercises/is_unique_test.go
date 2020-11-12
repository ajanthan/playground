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

func TestIsUniqueNoDS(t *testing.T) {
	assert.True(t, IsUniqueNoDS(""))
	assert.True(t, IsUniqueNoDS(" "))
	assert.True(t, IsUniqueNoDS("a"))
	assert.True(t, IsUniqueNoDS("trieyu"))
	assert.True(t, IsUniqueNoDS("4l1h3b"))
	assert.True(t, IsUniqueNoDS("fg?056*&("))

	assert.False(t, IsUniqueNoDS("aa"))
	assert.False(t, IsUniqueNoDS("a12344rtt"))
	assert.False(t, IsUniqueNoDS("*dgkoanmlw*"))
	assert.False(t, IsUniqueNoDS("*dgkoanmlw1234567890-=qwertyuioasdgjlxcvnm"))
}

func TestIsUniqueLowerCases(t *testing.T) {
	assert.True(t, IsUniqueLowerCases(""))
	assert.True(t, IsUniqueLowerCases(" "))
	assert.True(t, IsUniqueLowerCases("a"))
	assert.True(t, IsUniqueLowerCases("trieyu"))
	assert.False(t, IsUniqueLowerCases("abcda"))
	assert.False(t, IsUniqueLowerCases("abcdab"))
}
