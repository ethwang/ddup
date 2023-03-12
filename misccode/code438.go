package misccode

func FindAnagrams(s string, p string) []int {
	if len(s) < len(p) {
		return []int{}
	}
	res := []int{}
	pa := [26]int{}
	sa := [26]int{}
	for i := 0; i < len(p); i++ {
		pa[p[i]-'a']++
		sa[s[i]-'a']++
	}
	if pa == sa {
		res = append(res, 0)
	}
	for i := 1; i <= len(s)-len(p); i++ {
		sa = [26]int{}
		for j := i; j < i+len(p); j++ {
			sa[s[j]-'a']++
		}
		if pa == sa {
			res = append(res, i)
		}
	}
	return res
}
