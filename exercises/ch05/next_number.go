package ch05

//1010
//count=2
//msb=3
//0011
//1100
func NextNumber(b int) (int, int) {
	i, count, msb := 0, 0, 0
	for i < 32 {
		if !IsBitClear(b, i) {
			count++
			msb = i
		}
		i++
	}
	i = 0
	small := 0
	for i < count {
		small |= 1 << i
		i++
	}
	return small, small << (msb - count + 1)
}
