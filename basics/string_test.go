package basics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestString(t *testing.T) {
	//a string is a read-only slice of bytes.
	var name string
	assert.Equal(t, "", name)
	name += "ajanthan"
	assert.Equal(t, "ajanthan", name)
	assert.Equal(t, 8, len(name))
	assert.Equal(t, byte('a'), name[0])
	assert.Equal(t, "a", string(name[0]))
	assert.Equal(t, "a", name[0:1])
}
