package ch03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIntStack(t *testing.T) {
	is := &IntStack{}
	is.Push(3)
	is.Push(2)
	is.Push(4)
	is.Push(1)

	val, err := is.Min()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = is.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	val, err = is.Min()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	is.Push(2)

	val, err = is.Min()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	val, err = is.Min()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

}
