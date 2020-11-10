package basics

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMaps(t *testing.T) {
	var hashTable map[int]string
	assert.Equal(t, "", hashTable[2])
	//hashTable[1]="ajanthan" panic
	hashTable = make(map[int]string)
	zeroVal, ok := hashTable[0]
	assert.False(t, ok)
	assert.Equal(t, "", zeroVal)
	hashTable[1] = "ajanthan"
	name, ok := hashTable[1]
	assert.True(t, ok)
	assert.Equal(t, "ajanthan", name)
	assert.Equal(t, 1, len(hashTable))
}
