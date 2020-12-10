package ch05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsert(t *testing.T) {
	//654321
	//100000
	//  110
	n := 0b100000
	m := 0b110
	r := Insert(m, n, 4, 1)
	assert.Equal(t, 0b101100, r)
}
