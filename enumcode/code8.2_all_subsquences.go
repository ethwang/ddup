package enumcode

import "fmt"

func AllSubsquences(str string) {
	process([]byte(str), 0, "")
}

func process(chs []byte, i int, s string) {
	if i == len(chs) {
		fmt.Println(s)
		return
	}
	s += string(chs[i])
	process(chs, i+1, s)
	s = s[0 : len(s)-1]
	process(chs, i+1, s)
}
