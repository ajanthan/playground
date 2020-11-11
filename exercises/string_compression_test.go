package exercises

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicCompression(t *testing.T) {
	assert.Equal(t, "a2b1c5a3", BasicCompression("aabcccccaaa"))
	assert.Equal(t, "slddkkkvndk", BasicCompression("slddkkkvndk"))
	assert.Equal(t, "a3l3s3k3s1k1", BasicCompression("aaalllssskkksk"))
}
