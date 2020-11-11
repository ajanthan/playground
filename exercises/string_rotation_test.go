package exercises

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsRotated(t *testing.T) {
	assert.True(t, IsRotated("ajanthan", "janthana"))
	assert.True(t, IsRotated("aja123", "123aja"))
	assert.True(t, IsRotated("aja123", "23aja1"))
	assert.True(t, IsRotated("ajan", "ajan"))

	assert.False(t, IsRotated("ajan", "ajaa"))
	assert.False(t, IsRotated("january", "jamubry"))
	assert.False(t, IsRotated("bala", "mula"))
	assert.False(t, IsRotated("bala", "mula"))
	assert.False(t, IsRotated("bala", "ajan"))
}
