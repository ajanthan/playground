package basics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestArray(t *testing.T) {
	var intArray [3]int
	assert.NotNil(t, intArray)
	intArray[0] = 1
	intArray[2] = 3
	assert.Equal(t, 1, intArray[0])
	for index, i := range intArray {
		t.Log("intArray[", index, "]=", i)
	}
	prices := [...]int{23, 56, 78}
	assert.Equal(t, 78, prices[2])

}
