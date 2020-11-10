package basics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSlices(t *testing.T) {
	names := make([]string, 0)
	names = append(names, "ajanthan")
	assert.Equal(t, "ajanthan", names[0])
	address := []string{"Sunnyvale", "Mountain View", "San Jose"}

	for index, i := range address {
		t.Log("address[", index, "]=", i)
	}

	assert.Equal(t, "Mountain View", address[1])
	assert.Equal(t, 3, len(address))
	//remove "Mountain View"
	address = append(address[:1], address[2:]...)
	assert.Equal(t, 2, len(address))
	assert.Equal(t, "San Jose", address[1])
}
