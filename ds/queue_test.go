package ds

import "testing"
import "github.com/stretchr/testify/assert"

func TestQueue(t *testing.T) {
	queue:=&Queue{}
	queue.Put("test0")
	queue.Put("test1")
	queue.Put("test2")

	s,err:=queue.Get()
	assert.NoError(t,err)
	assert.Equal(t,"test0",s)

	s,err=queue.Get()
	assert.NoError(t,err)
	assert.Equal(t,"test1",s)

	s,err=queue.Get()
	assert.NoError(t,err)
	assert.Equal(t,"test2",s)

	s,err=queue.Get()
	assert.EqualError(t,err,"empty queue")
	assert.Equal(t,"",s)


}