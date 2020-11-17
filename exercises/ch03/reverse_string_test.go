package ch03

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"testing"
)

func TestReverse(t *testing.T) {
	cases := []struct {
		f func(s string) string
		d string
	}{
		{
			f: ReverseV1,
			d: "Recursive",
		},
		{
			f: ReverseV0,
			d: "Iterative",
		},
	}
	for i, c := range cases {
		s := ""
		assert.Equal(t, "", c.f(s))
		s = "data structure"
		assert.Equal(t, "erutcurts atad", c.f(s), fmt.Sprintf("Case %d: %s", i, c.d))
		s = "bumble bee"
		assert.Equal(t, s, c.f(c.f(s)), fmt.Sprintf("Case %d: %s", i, c.d))
	}
}
