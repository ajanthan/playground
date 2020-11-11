package exercises

import (
	"strconv"
	"strings"
)

func BasicCompression(s string) string { //aabccaa-->a2b1c2
	var builder strings.Builder
	prevChar := s[0]
	count := 1
	for i := 1; i < len(s); i++ {
		if s[i] == prevChar {
			count++
		} else {
			builder.WriteString(string(prevChar) + strconv.Itoa(count))
			if builder.Len() >= len(s) {
				return s
			}
			prevChar = s[i]
			count = 1
		}
	}
	builder.WriteString(string(prevChar) + strconv.Itoa(count))
	return builder.String()
}
