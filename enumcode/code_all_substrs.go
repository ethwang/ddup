package enumcode

// abc 子串   a,b,c,ab,bc,abc
func AllSubStrs(s string) []string {

	ss := make([]string, 0)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			ss = append(ss, s[i:j])
		}
	}
	return ss
}
