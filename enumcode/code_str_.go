package enumcode

func numToChars(chs []byte, new *int, cnt int) {
	cnts := make([]int, 0)
	cnt_ := cnt % 10
	for ; cnt_ != 0; cnt_ = cnt % 10 {
		cnts = append(cnts, cnt_)
		cnt = cnt / 10
	}
	for i := len(cnts) - 1; i >= 0; i-- {
		chs[*new] = byte(cnts[i] + '0')
		(*new)++
	}
}

// 字符串中相同字符数量统计，空间复杂度O(1)
func SameCharsCount(str string) string {
	// 1.字符串转化为字符数组；
	// 2.遍历字符数组，统计相同字符个数并找到相同字符下的最后一个字符；
	// 3.将字符个数重写回字符串最后一个字符位置，并删除多余字符;

	chs := []byte(str)

	cnt := 0
	i := 0
	new := 0

	for ; i < len(chs); i++ {
		if i > 0 && chs[i] != chs[i-1] {
			chs[new] = chs[i-1]
			new++
			if cnt > 1 {
				numToChars(chs, &new, cnt)
			}
			cnt = 1
		} else {
			cnt++
		}
	}
	if cnt > 1 {
		chs[new] = chs[i-1]
		new++
		numToChars(chs, &new, cnt)
	}

	return string(chs[:new])
}
