package misccode

func Rotate1(matrix [][]int) {
	n := len(matrix)
	newMatrix := make([][]int, n)
	for i := 0; i < n; i++ {
		newMatrix[i] = make([]int, n)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			newMatrix[i][j] = matrix[j][n-i-1]
		}
	}
	copy(matrix, newMatrix)
}

func Rotate2(matrix [][]int) {
	n := len(matrix)

	if n%2 == 0 {

		for i := 0; i < (n / 2); i++ {
			for j := 0; j < (n / 2); j++ {

				tmp := matrix[i][j]
				matrix[i][j] = matrix[n-j-1][i]
				matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
				matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
				matrix[j][n-i-1] = tmp
			}
		}
	} else {

		for i := 0; i < (n-1)/2; i++ {
			for j := 0; j < (n+1)/2; j++ {
				tmp := matrix[i][j]
				matrix[i][j] = matrix[n-j-1][i]
				matrix[n-j-1][i] = matrix[n-i-1][n-j-1]
				matrix[n-i-1][n-j-1] = matrix[j][n-i-1]
				matrix[j][n-i-1] = tmp

			}
		}
	}
}
