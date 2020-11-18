package ch03

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSubStack(t *testing.T) {
	ss := &SubStack{}
	ss.Push(1)
	val, err := ss.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	ss.Push(2)
	ss.Push(1)
	val, err = ss.PopAt(0)
	assert.NoError(t, err)
	assert.Equal(t, 1, val)

	_, err = ss.PopAt(1)
	assert.Error(t, err)

	val, err = ss.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 2, val)

	for i := 0; i < 35; i++ {
		ss.Push(i)
	}

	val, err = ss.Pop()
	assert.NoError(t, err)
	assert.Equal(t, 34, val)

	_, err = ss.PopAt(0)
	assert.NoError(t, err)
	_, err = ss.PopAt(1)
	assert.NoError(t, err)
	_, err = ss.PopAt(2)
	assert.NoError(t, err)

	_, err = ss.PopAt(3)
	assert.NoError(t, err)
	ss.PopAt(3)
	ss.PopAt(3)
	_, err = ss.PopAt(3)
	assert.Error(t, err)

}
