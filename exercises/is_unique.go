package exercises

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
