package basics

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestConcatStrings(t *testing.T) {
	words := []string{"Golang ", "is", " awesome"}
	var buffer strings.Builder
	for _, word := range words {
		buffer.WriteString(word)
	}
	sentence := buffer.String()
	assert.Equal(t, "Golang is awesome", sentence)
}
