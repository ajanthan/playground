package ch05

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToString(t *testing.T) {

	r, err := ToString(0.5)
	assert.NoError(t, err)
	assert.Equal(t, "1", r)

	r, err = ToString(0.25)
	assert.NoError(t, err)
	assert.Equal(t, "01", r)

	r, err = ToString(0.75)
	assert.NoError(t, err)
	assert.Equal(t, "11", r)

	r, err = ToString(0.125)
	assert.NoError(t, err)
	assert.Equal(t, "001", r)

	r, err = ToString(0.875)
	assert.NoError(t, err)
	assert.Equal(t, "111", r)

	r, err = ToString(0.59375)
	assert.NoError(t, err)
	assert.Equal(t, "10011", r)

	r, err = ToString(0.50000000069849193096)
	assert.NoError(t, err)
	assert.Equal(t, "10000000000000000000000000000011", r)

	r, err = ToString(0.01562500069849193096)
	assert.NoError(t, err)
	assert.Equal(t, "00000100000000000000000000000011", r)

	r, err = ToString(0.00000000069849193096)
	assert.Error(t, err)
}

//0.75*2=1.50
//0.50*2=1
//0--->11=0.5+0.25
