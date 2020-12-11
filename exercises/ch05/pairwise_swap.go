package ch05

//10->01
func SwapOddEvenBits(b int) int {
	even := b << 1 //100
	odd := b >> 1  //001
	i, r := 0, 0
	for i < 32 {
		if !IsBitClear(odd, i) {
			r |= 1 << i
		}
		i += 2
	}
	i = 1
	for i < 32 {
		if !IsBitClear(even, i) {
			r |= 1 << i
		}
		i += 2
	}
	return r
}
