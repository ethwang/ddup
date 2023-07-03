package misccode

func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}
	if s != "" && p == "" {
		return false
	}
	if s == "" && p != "" {
		if len(p) == 2 && p[1] == '*' {
			return true
		}
		return false
	}
	dp := make([][]bool, len(s)+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]bool, len(p))
	}
	dp[0][0] = true

	// dp[i][j]表示s[0,i-1]和p[0,j-1]是否匹配
	for i := 1; i <= len(s); i++ {
		for j := 1; j <= len(p); j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i-1][j-1]
			}
			if s[i-1] != p[j-1] && p[j-1] == '*' {
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i-1][j-2] || dp[i][j-2] || dp[i-1][j]
				} else {
					dp[i][j] = dp[i][j-2]
				}
			}
		}
	}
	return dp[len(s)][len(p)]
}
