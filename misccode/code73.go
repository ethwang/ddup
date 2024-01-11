package misccode

func SetZeroes(matrix [][]int) {
	row, col := false, false
	rowLen, colLen := len(matrix), len(matrix[0])

	for i := 0; i < rowLen; i++ {
		if matrix[i][0] == 0 {
			col = true
			break
		}
	}
	for j := 0; j < colLen; j++ {
		if matrix[0][j] == 0 {
			row = true
			break
		}
	}
	for i := 1; i < rowLen; i++ {
		for j := 1; j < colLen; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	for i := 0; i < rowLen; i++ {
		if matrix[i][0] == 0 {
			for j := 1; j < colLen; j++ {
				matrix[i][j] = 0
			}
		}
	}
	for j := 0; j < colLen; j++ {
		if matrix[0][j] == 0 {
			for i := 1; i < rowLen; i++ {
				matrix[i][j] = 0
			}
		}
	}
	if col {
		for i := 0; i < rowLen; i++ {
			matrix[i][0] = 0
		}
	}
	if row {
		for j := 0; j < colLen; j++ {
			matrix[0][j] = 0
		}
	}
}
