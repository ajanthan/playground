package ch01

import "strings"

func URLify(s string, trueLength int) string {
	var buffer strings.Builder
	trueCount := 0
	for i := 0; i < len(s); i++ {
		if s[i] == byte(' ') {
			if trueCount == 0 {
				continue
			} else if trueCount == trueLength {
				break
			}
			trueCount++
			buffer.WriteString("%20")
		} else {
			trueCount++
			buffer.WriteByte(s[i])
		}
	}
	return buffer.String()
}
