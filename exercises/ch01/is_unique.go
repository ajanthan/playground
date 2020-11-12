package ch01

func IsUnique(s string) bool {
	ht := make(map[byte]string)
	for i := 0; i < len(s); i++ {
		_, ok := ht[s[i]]
		if ok {
			return false
		} else {
			ht[s[i]] = string(s[i])
		}
	}
	return true
}

func IsUniqueNoDS(s string) bool {
	if len(s) > 128 {
		return false
	}
	var chart [128]bool
	for i := 0; i < len(s); i++ {
		if chart[s[i]] {
			return false
		} else {
			chart[s[i]] = true
		}
	}
	return true
}
func IsUniqueLowerCases(s string) bool {
	checker := 0
	for i := 0; i < len(s); i++ {
		index := s[i] - 'a'
		if (checker & (1 << index)) > 0 {
			return false
		} else {
			checker |= 1 << index
		}
	}
	return true
}
