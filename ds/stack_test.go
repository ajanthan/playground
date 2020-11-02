package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := NewStack()
	stack.Push("test0")
	stack.Push("test1")
	stack.Push("test2")
	val, err := stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "test2", val)

	val, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "test1", val)

	val, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "test0", val)

	val, err = stack.Pop()
	assert.EqualError(t, err, "empty stack")
	assert.Equal(t, "", val)

	stack.Push("test4")
	stack.Push("test5")
	stack.Clear()

	val, err = stack.Pop()
	assert.EqualError(t, err, "empty stack")
	assert.Equal(t, "", val)
}
