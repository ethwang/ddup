package misccode

func MaximalSquare(matrix [][]byte) int {
	dp := make([][]int, len(matrix))
	for i := 0; i < len(matrix); i++ {
		dp[i] = make([]int, len(matrix[i]))
	}

	for i := 0; i < len(matrix); i++ {
		if matrix[i][0] == '1' {
			dp[i][0] = 1
		}
	}
	for j := 0; j < len(matrix[0]); j++ {
		if matrix[0][j] == '1' {
			dp[0][j] = 1
		}
	}
	res := 0

	for i := 1; i < len(matrix); i++ {
		for j := 1; j < len(matrix[i]); j++ {
			if matrix[i][j] == 1 {
				dp[i][j] = min(dp[i-1][j], min(dp[i][j-1], dp[i-1][j-1])) + 1
				if dp[i][j] > res {
					res = dp[i][j]
				}
			}
		}
	}
	return res * res
}
