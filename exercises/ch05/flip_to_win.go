package ch05

//11001
//i=2
func IsBitClear(b, i int) bool {
	return (b & (1 << i)) == 0
}

func LongestOnesAfterAFlip(b int) int {
	i, count, max, start := 0, 0, 0, 0
	cleared := false
	for i < 32 {
		i = start
		for i < 32 {
			if IsBitClear(b, i) {
				if !cleared {
					cleared = true
					start = i + 1
				} else {
					if count > max {
						max = count
					}
					count = 0
					cleared = false
					break
				}
			}
			count++
			i++
		}
	}
	return max
}
