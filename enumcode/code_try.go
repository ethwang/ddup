package enumcode

import "fmt"

func Try(s string) {
	chs := []byte(s)

	fmt.Println(process5(chs, 0, len(chs)))
}
func process5(chs []byte, i, n int) int {
	if i == n {
		return 1
	}
	if chs[i] == '0' {
		return 0
	}
	if chs[i] == '1' {
		res := process5(chs, i+1, n)
		res += process5(chs, i+2, n)
		return res
	}
	if chs[i] == '2' {

		res := process5(chs, i+1, n)
		if chs[i+1] >= '0' && chs[i+1] <= '6' {

			res += process5(chs, i+2, n)
		}
	}
	return process5(chs, i+1, n)
}

func process6(chs []byte, i, n int, res string) {

	if i == n {
		return
	}
	for j := 0; j < len(chs); j++ {
		res += string(chs[j])
		process6(chs, i+1, n, res)

	}
}
