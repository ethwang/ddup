package misccode

func WordBreak(s string, wordDict []string) bool {
	m := make(map[string]bool)
	for _, v := range wordDict {
		m[v] = true
	}
	emeo := make(map[int]bool)
	return process111(0, m, emeo, s)
	//return process11(0, m, s)
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

// s="abcd" worldDict:=["a","b","c","ab","bc"]
func process111(start int, m map[string]bool, emeo map[int]bool, s string) bool {
	if start == len(s) {
		return true
	}
	if v, ok := emeo[start]; ok {
		return v
	}
	for i := start + 1; i <= len(s); i++ {
		str := s[start:i]
		if !m[str] {
			continue
		}
		res := process111(i, m, emeo, s)
		if res {
			emeo[start] = true
			return true
		}
	}
	emeo[start] = false
	return false
}
