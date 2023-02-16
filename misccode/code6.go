package misccode

import (
	"fmt"
	"strconv"
)

func LetterCombinations(digits string) []string {
	ress := make([]string, 0)
	chs := []byte(digits)
	var res string

	if digits == "" {
		return ress
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}

	process(0, len(digits), chs, m, res, &ress)
	return ress
}
func process(i, n int, chs []byte, m map[byte][]string, res string, ress *[]string) {

	if i == n {
		*ress = append(*ress, res)
		return
	}
	if v, ok := m[chs[i]]; ok {
		for j := 0; j < len(v); j++ {
			res += string(v[j])
			process(i+1, n, chs, m, res, ress)
			res = res[:len(res)-1]
		}
	}
}

func LetterCombinations2(digits string) []string {
	var rs = map[int][]string{}
	rs[2] = []string{"a", "b", "c"}
	rs[3] = []string{"d", "e", "f"}
	rs[4] = []string{"g", "h", "i"}
	rs[5] = []string{"j", "k", "l"}
	rs[6] = []string{"m", "n", "o"}
	rs[7] = []string{"p", "q", "r", "s"}
	rs[8] = []string{"v", "t", "u"}
	rs[9] = []string{"w", "x", "y", "z"}

	var rst = map[string]string{}

	if len(digits) == 0 {
		return []string{}
	}

	// 遍历传入的数字串
	for _, key := range []rune(digits) {
		tempKey, _ := strconv.Atoi(fmt.Sprintf("%c", key))

		if val, ok := rs[tempKey]; ok {
			if len(rst) == 0 {
				for _, firstVal := range val {
					rst[firstVal] = firstVal
				}
				//fmt.Println("第一次，", rst)
				// a, b, c
			} else {
				// 开始时：a b c; 第二轮:d e
				var tempMap = make(map[string]string)
				for _, rstVal := range rst {
					tempMap[rstVal] = rstVal
				}

				rst = map[string]string{}

				for _, str := range val { // d e
					for _, rstVal := range tempMap {
						kk := fmt.Sprintf("%v%v", rstVal, str)
						rst[kk] = kk
					}
				}
			}
		}
	}

	result := []string{}
	for _, rstVal := range rst {
		result = append(result, rstVal)
	}

	return result
}
