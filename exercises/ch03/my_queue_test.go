package ch03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMyQueue(t *testing.T) {
	myQueue := NewMyQueue()

	assert.True(t, myQueue.IsEmpty())

	myQueue.Add("a")
	myQueue.Add("b")
	myQueue.Add("c")
	assert.False(t, myQueue.IsEmpty())

	val, err := myQueue.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "a", val)

	val, err = myQueue.Remove()
	assert.NoError(t, err)
	assert.Equal(t, "a", val)

	val, err = myQueue.Remove()
	assert.NoError(t, err)
	assert.Equal(t, "b", val)

	myQueue.Add("a")
	myQueue.Add("b")

	val, err = myQueue.Remove()
	assert.NoError(t, err)
	assert.Equal(t, "c", val)

	val, err = myQueue.Remove()
	assert.NoError(t, err)
	assert.Equal(t, "a", val)

	val, err = myQueue.Remove()
	assert.NoError(t, err)
	assert.Equal(t, "b", val)

	_, err = myQueue.Remove()
	assert.Error(t, err)

	assert.True(t, myQueue.IsEmpty())

}
