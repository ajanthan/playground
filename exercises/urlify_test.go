package exercises

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURLify(t *testing.T) {
	assert.Equal(t, "Mr%20John%20Smith", URLify("Mr John Smith", 13))
	assert.Equal(t, "Mr%20John%20Smith", URLify("Mr John Smith ", 13))
	assert.Equal(t, "Mr%20John%20Smith", URLify(" Mr John Smith ", 13))
	assert.Equal(t, "Mr%20John%20Smith", URLify("  Mr John Smith ", 13))
	assert.Equal(t, "Mr%20John%20Smith", URLify("  Mr John Smith  ", 13))
}
