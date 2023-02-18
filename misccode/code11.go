package misccode

func WordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true
	}
	return process11(0, m, s)
}
func process11(start int, m map[string]bool, s string) bool {
	if start == len(s) {
		return true
	}
	for i := start + 1; i <= len(s); i++ {
		str := s[start:i]
		if _, ok := m[str]; ok && process11(i, m, s) {
			return true
		}
	}
	return false
}
