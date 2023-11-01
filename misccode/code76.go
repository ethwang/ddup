package misccode

import "math"

func MinWindow(s string, t string) string {
	ans := ""
	lmin := math.MaxInt
	rt, st := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		rt[t[i]]++
	}
	for l, r := 0, 0; r < len(s); r++ {
		if rt[s[r]] > 0 {
			st[s[r]]++
		}
		for check(rt, st) {
			if lmin > r-l+1 {
				lmin = r - l + 1
				ans = s[l : l+lmin]
			}
			if st[s[l]] > 0 {
				st[s[l]]--
			}
			l++
		}
	}
	return ans
}
func check(rt, st map[byte]int) bool {
	for k, v := range rt {
		if v > st[k] {
			return false
		}
	}
	return true
}
