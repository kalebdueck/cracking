package chapter1

func ZeroMatrix(matrix [][]int) [][]int {
	var zeroRows []int
	var zeroCols []int

	for row, cols := range matrix {
		for col, val := range cols {
			if val == 0 {
				zeroCols = append(zeroCols, col)
				zeroRows = append(zeroRows, row)
			}
		}
	}

	colLength := len(matrix[0])

	for _, row := range zeroRows {
		matrix[row] = make([]int, colLength)
	}

	for _, col := range zeroCols {
		for row, _ := range matrix {
			matrix[row][col] = 0
		}
	}

	return matrix
}
