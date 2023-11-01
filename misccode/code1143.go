package misccode

func longestCommonSubsequence(text1 string, text2 string) int {
	m, n := len(text1)+1, len(text2)+1
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i, t1c := range text1 {
		for j, t2c := range text2 {
			if t1c == t2c {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = max(dp[i+1][j], dp[i][j+1])
			}
		}
	}
	return dp[len(text1)][len(text2)]
}
