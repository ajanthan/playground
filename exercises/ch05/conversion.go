package ch05

//a=11101
//b=01111
func NumberOfBitFlipToConvert(a, b int) int {
	i, count := 0, 0
	for i < 32 {
		if IsBitClear(a, i) != IsBitClear(b, i) {
			count++
		}
		i++
	}
	return count
}
