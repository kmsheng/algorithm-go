package question

import "fmt"

func ZeroMatrix(matrix [][]int) [][]int {
	size := len(matrix)
	m := map[string]bool{}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if matrix[i][j] == 0 {
				m[fmt.Sprintf("%i:i", i)] = true
				m[fmt.Sprintf("%i:j", j)] = true
			}
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if m[fmt.Sprintf("%i:i", i)] || m[fmt.Sprintf("%i:j", j)] {
				matrix[i][j] = 0
			}
		}
	}
	return matrix
}
