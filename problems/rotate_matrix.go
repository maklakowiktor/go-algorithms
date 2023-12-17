package problems

func RotateMatrix(matrix [][]int) [][]int {
	// Get matrix sizes
	rows := len(matrix)
	cols := len(matrix[0])

	// Create new matrix with transposed elements
	rotatedMatrix := make([][]int, cols)
	for i := 0; i < cols; i++ {
		rotatedMatrix[i] = make([]int, rows)
	}

	// Transpose matrix's elements
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			rotatedMatrix[j][i] = matrix[i][j]
		}
	}

	// Reverse each row
	for i := 0; i < rows; i++ {
		for j := 0; j < cols/2; j++ {
			rotatedMatrix[i][j], rotatedMatrix[i][len(matrix[i])-1-j] = rotatedMatrix[i][len(rotatedMatrix[i])-1-j], rotatedMatrix[i][j]
		}
	}

	return rotatedMatrix
}
