package ch01

import (
	"strings"
)

func IsRotated(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	i, j := 0, 0
	for i < len(s1) && j < len(s2) {
		if s1[i] == s2[j] {
			i++
			j++
		} else {
			j++
			i = 0
		}
	}
	if i == 0 {
		return false
	}
	return strings.Index(s2, s1[i:]+s1[:i]) > -1
}
