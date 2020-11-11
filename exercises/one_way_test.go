package exercises

import "strings"

func IsOneAway(s1, s2 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	if s1 == s2 {
		return true
	}
	var newStr strings.Builder
	if l1 > l2 {
		//remove
		for i := 0; i < len(s1); i++ {
			newStr.WriteString(s1[:i])
			newStr.WriteString(s1[i+1:])
			if s2 == newStr.String() {
				return true
			}
		}
	} else if l1 < l2 {
		//insert
		for i := 0; i < len(s2); i++ {
			newStr.WriteString(s2[:i])
			newStr.WriteString(s2[i+1:])
			if s1 == newStr.String() {
				return true
			}
		}
	} else {
		//change
		chnaged := 0
		for i := 0; i < len(s2); i++ {
			if s1[i] != s2[i] {
				chnaged++
			}
		}
		if chnaged == 1 {
			return true
		}
	}
	return false
}
