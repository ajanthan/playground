package exercises

func RotateImage90(image [][]int) [][]int {
	newImage := make([][]int, len(image))
	for i := range image {
		newImage[i] = make([]int, len(image))
	}
	for i, r := range image {
		for j := range r {
			newImage[j][len(image)-1-i] = image[i][j]
		}
	}
	return newImage
}
