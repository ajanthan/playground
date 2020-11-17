package ds

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStack(t *testing.T) {
	stack := &Stack{}
	assert.True(t, stack.IsEmpty())
	stack.Push("test0")
	assert.False(t, stack.IsEmpty())
	stack.Push("test1")
	val, err := stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "test1", val)
	stack.Push("test2")
	val, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "test2", val)

	val, err = stack.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "test1", val)

	val, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "test1", val)

	val, err = stack.Pop()
	assert.NoError(t, err)
	assert.Equal(t, "test0", val)

	val, err = stack.Pop()
	assert.EqualError(t, err, "empty stack")
	assert.Equal(t, "", val)
	assert.True(t, stack.IsEmpty())

	stack.Push("test4")
	stack.Push("test5")
	stack.Clear()
	assert.True(t, stack.IsEmpty())

	val, err = stack.Pop()
	assert.EqualError(t, err, "empty stack")
	assert.Equal(t, "", val)
}
