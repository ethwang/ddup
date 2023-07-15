package misccode

func MaxCoins(nums []int) int {
	tmpNums := make([]int, len(nums)+2)
	for i := 1; i <= len(nums); i++ {
		tmpNums[i] = nums[i-1]
	}
	tmpNums[0], tmpNums[len(nums)+1] = 1, 1
	dp := make([][]int, len(nums)+2)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, len(nums)+2)
	}

	// 区间长度
	for i := 3; i <= len(tmpNums); i++ {
		// 区间起点
		for j := 0; j <= len(tmpNums)-i; j++ {
			for k := j + 1; k < j+i-1; k++ {
				dp[j][j+i-1] = max(dp[j][j+i-1], dp[j][k]+dp[k][j+i-1]+tmpNums[k]*tmpNums[j]*tmpNums[j+i-1])
			}
		}
	}
	return dp[0][len(tmpNums)]
}
