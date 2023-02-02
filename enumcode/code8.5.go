package enumcode

func process3(chs []byte, i int) int {
	if i == len(chs) {
		return 1
	}
	if chs[i] == '0' {
		return 0
	}
	res := 0
	if chs[i] == '1' {
		res = process3(chs, i+1)
		if i+1 < len(chs) {
			res += process3(chs, i+2)
		}
	}
	if chs[i] == '2' {
		res = process3(chs, i+1)
		if i+1 < len(chs) && chs[i+1] >= '0' && chs[i+1] <= '6' {
			res += process3(chs, i+2)
		}
	}
	res += process3(chs, i+1)
	return res
}
