package ch02

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"playground/ds"
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	cases := []struct {
		f func(l ds.LinkedList) bool
		d string
	}{
		{
			f: IsPalindromeV0,
			d: "reverse equal method",
		},
		{
			f: IsPalindromeV1,
			d: "slow runner fast runner",
		},
		{
			f: IsPalindromeV2,
			d: "recursion",
		},
	}
	for i, c := range cases {
		l1 := &ds.LinkedList{}
		l1.AddNode(1)
		l1.AddNode(2)
		l1.AddNode(3)
		l1.AddNode(2)
		l1.AddNode(1)
		assert.True(t, c.f(*l1), fmt.Sprintf("Case %d: %s", i, c.d))
		l2 := &ds.LinkedList{}
		l2.AddNode(1)
		l2.AddNode(2)
		l2.AddNode(3)
		l2.AddNode(3)
		l2.AddNode(2)
		l2.AddNode(1)
		assert.True(t, c.f(*l2), fmt.Sprintf("Case %d: %s", i, c.d))
		l1.DeleteNode(3)
		assert.True(t, c.f(*l1), fmt.Sprintf("Case %d: %s", i, c.d))
		l2.DeleteNode(1)
		assert.False(t, c.f(*l2), fmt.Sprintf("Case %d: %s", i, c.d))
	}

}
