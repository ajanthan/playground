package ch06

func FindSteps(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	} else if n == 3 {
		return 3
	}
	possibleWays := 3
	for x := 0; x < n; x++ {
		for y := 0; y < n/2; y++ {
			for z := 0; z < n/3; z++ {
				steps := x + 2*y + 3*z
				if steps == n {
					possibleWays++
				}
			}
		}
	}
	return possibleWays
}
