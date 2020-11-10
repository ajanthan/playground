package exercises

func CheckPermutation(s1, s2 string) bool {
	if s1 == s2 {
		return true
	} else if len(s1) != len(s2) {
		return false
	}
	ht := make(map[byte][]string)
	for i := 0; i < len(s1); i++ {
		ht[s1[i]] = append(ht[s1[i]], string(s1[i]))
	}
	for i := 0; i < len(s2); i++ {
		ss, found := ht[s2[i]]
		if found {
			l := len(ss)
			if l == 1 {
				delete(ht, s2[i])
			} else {
				ss = ss[:l-1]
				ht[s2[i]] = ss
			}
		} else {
			return false
		}
	}
	if len(ht) == 0 {
		return true
	}
	return false
}
