package ds

import "testing"
import "github.com/stretchr/testify/assert"

func TestQueue(t *testing.T) {
	queue := &Queue{}
	assert.True(t, queue.IsEmpty())
	queue.Add("test0")
	assert.False(t, queue.IsEmpty())
	queue.Add("test1")
	s, err := queue.Peek()
	assert.NoError(t, err)
	assert.Equal(t, "test0", s)

	queue.Add("test2")

	s, err = queue.Remove()
	assert.NoError(t, err)
	assert.Equal(t, "test0", s)

	s, err = queue.Remove()
	assert.NoError(t, err)
	assert.Equal(t, "test1", s)

	s, err = queue.Remove()
	assert.NoError(t, err)
	assert.Equal(t, "test2", s)

	_, err = queue.Peek()
	assert.Error(t, err)

	s, err = queue.Remove()
	assert.EqualError(t, err, "empty queue")
	assert.Equal(t, "", s)
	assert.True(t, queue.IsEmpty())

}
