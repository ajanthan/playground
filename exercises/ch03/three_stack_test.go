package ch03

import (
	"fmt"
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestThreeStack_Push1(t *testing.T) {
	ts := NewThreeStack()
	assert.True(t, ts.IsEmpty1())
	assert.True(t, ts.IsEmpty2())
	assert.True(t, ts.IsEmpty3())
	ts.Push1(fmt.Sprintf("test1[%d]", 0))
	ts.Push2(fmt.Sprintf("test2[%d]", 0))
	ts.Push3(fmt.Sprintf("test3[%d]", 0))

	val, err := ts.Peek1()
	assert.NoError(t, err)
	assert.Equal(t, "test1[0]", val)

	val, err = ts.Peek2()
	assert.NoError(t, err)
	assert.Equal(t, "test2[0]", val)

	val, err = ts.Peek3()
	assert.NoError(t, err)
	assert.Equal(t, "test3[0]", val)

	val, err = ts.Pop1()
	assert.NoError(t, err)
	assert.Equal(t, "test1[0]", val)

	val, err = ts.Pop2()
	assert.NoError(t, err)
	assert.Equal(t, "test2[0]", val)

	val, err = ts.Pop3()
	assert.NoError(t, err)
	assert.Equal(t, "test3[0]", val)

	for i := 0; i < 11; i++ {
		ts.Push1(fmt.Sprintf("test1[%d]", i))
		ts.Push2(fmt.Sprintf("test2[%d]", i))
		ts.Push3(fmt.Sprintf("test3[%d]", i))
	}

	ts.Push1(fmt.Sprintf("test1[%d]", 11))
	val, err = ts.Pop1()
	assert.NoError(t, err)
	assert.Equal(t, "test1[11]", val)

	ts.Push2(fmt.Sprintf("test2[%d]", 11))
	val, err = ts.Pop2()
	assert.NoError(t, err)
	assert.Equal(t, "test2[11]", val)

	ts.Push3(fmt.Sprintf("test3[%d]", 11))
	val, err = ts.Pop3()
	assert.NoError(t, err)
	assert.Equal(t, "test3[11]", val)

}
