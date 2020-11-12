package ch01

func ZeroMatrix(mn [][]int) [][]int {
	newMatrix := make([][]int, len(mn))
	for i := range mn {
		newMatrix[i] = append(newMatrix[i], mn[i]...)
	}
	for i, r := range mn {
		for j := range r {
			if mn[i][j] == 0 {
				for k := range r {
					newMatrix[i][k] = 0
				}
				for l := range mn {
					newMatrix[l][j] = 0
				}
			}
		}
	}
	return newMatrix
}
