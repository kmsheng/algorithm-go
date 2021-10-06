package question

func RotateMatrix(matrix [][]int) [][]int {
	size := len(matrix)
	iEnd := size / 2
	for i := 0; i < iEnd; i++ {
		jEnd := size - i - 1
		for j := i; j < jEnd; j++ {
			// bottom right
			tmp := matrix[j][jEnd]
			matrix[j][jEnd] = matrix[i][j]

			// bottom left
			tmp2 := matrix[jEnd][jEnd - j + i]
			matrix[jEnd][jEnd - j + i] = tmp

			// top left
			tmp3 := matrix[jEnd - j + i][i]
			matrix[jEnd - j + i][i] = tmp2

			// top right
			matrix[i][j] = tmp3
		}
	}
	return matrix
}
