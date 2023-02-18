package misccode

func LongestDupSubstring(s string) string {
	m := make(map[string]int)
	maxLen := 0
	maxLenStr := ""
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			if j-i+1 < maxLen {
				continue
			}
			m[s[i:j]]++
			if m[s[i:j]] > 1 {
				if j-i+1 > maxLen {

					maxLen = j - i + 1
					maxLenStr = s[i:j]
				}
			}
		}
	}
	return maxLenStr
}
