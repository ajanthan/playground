package exercises

import (
	"regexp/syntax"
	"strings"
)

func IsPalindromePermutation(s string) bool {
	s = strings.ToLower(s) //ab ba
	ht := make(map[byte]int)
	for i := 0; i < len(s); i++ { //a   b  ' ' b
		if syntax.IsWordChar(rune(s[i])) {
			_, ok := ht[s[i]] //true
			if ok {
				delete(ht, s[i]) //del b
			} else {
				ht[s[i]] = 1 //a->1 b->1
			}
		}
	}
	if len(ht) > 1 {
		return false
	}
	return true
}
