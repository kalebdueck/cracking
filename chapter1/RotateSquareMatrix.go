package chapter1

func RotateSquareMatrix(matrix [][]int) [][]int {

	length := len(matrix)

	//Initial loop that dives through the "edges"
	for layer := 0; layer < length/2; layer++ {

		first := layer
		last := length - 1 - layer

		for i := first; i < last; i++ {
			offset := i - first
			//Temp var for swapping
			top := matrix[first][i]

			//left -> top
			matrix[first][i] = matrix[last-offset][first]

			//bottom -> left
			matrix[last-offset][first] = matrix[last][last-offset]

			//right -> bottom
			matrix[last][last-offset] = matrix[i][last]

			//top(temp value) -> right
			matrix[i][last] = top
		}

	}

	return matrix
}
