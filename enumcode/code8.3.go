package enumcode

import "fmt"

func Permutations(s string) {
	visit := make(map[byte]bool)
	process2([]byte(s), 0, "", visit)
}

func process2(chs []byte, i int, res string, visit map[byte]bool) {
	if i == len(chs) {
		fmt.Println(res)
		return
	}
	for j := 0; j < len(chs); j++ {
		if !visit[chs[j]] {
			visit[chs[j]] = true
			res += string(chs[j])
			process2(chs, i+1, res, visit)
			res = res[0 : len(res)-1]
			visit[chs[j]] = false
		}
	}
}
