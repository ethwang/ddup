package greedycode

import "sort"

// 自定义排序需要一个类型
type SS []string

func (ss SS) Len() int {
	return len(ss)
}
func (ss SS) Swap(i, j int) {
	ss[i], ss[j] = ss[j], ss[i]
}
func (ss SS) Less(i, j int) bool {
	// 字符串按照字典序是可以比较的
	return ss[i]+ss[j] < ss[j]+ss[i]
}

func MinDicOrder(ss []string) (s string) {

	// 局部最小->全局最小
	// 字符串数组从小到大排序
	// sort.Strings(ss)
	sort.Sort(SS(ss))

	for _, v := range ss {
		s += v
	}

	return
}
