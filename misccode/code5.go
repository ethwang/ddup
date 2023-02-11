package misccode

import "c1/linklistcode"

func LengthOfLongestSubstring(s string) int {

	chs := []byte(s)
	m := make(map[byte]int)
	maxLength := 0
	for i, j := 0, 0; i < len(chs) && j < len(chs); j++ {
		m[chs[j]]++
		if v, ok := m[chs[j]]; ok && v > 1 {
			for ; m[chs[j]] > 1 && i < j; i++ {
				m[chs[i]]--
			}
		} else {
			if j-i+1 > maxLength {
				maxLength = j - i + 1
			}
		}

	}
	return maxLength

}
func AddTwoNumbers(l1 *linklistcode.Node, l2 *linklistcode.Node) *linklistcode.Node {

	n := &linklistcode.Node{}
	res := n
	t1 := 0
	for l1 != nil || l2 != nil {
		t2 := 0
		l1v, l2v := 0, 0
		if l1 != nil {
			l1v = l1.Value
		}
		if l2 != nil {
			l2v = l2.Value
		}
		v := l1v + l2v + t1
		t1 = v / 10
		t2 = v % 10

		n.Value = t2
		if (l1 != nil && l1.Next != nil) || (l2 != nil && l2.Next != nil) {
			n.Next = &linklistcode.Node{}
			n = n.Next
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if t1 > 0 {
		n.Next = &linklistcode.Node{}
		n = n.Next
		n.Value = t1
	}
	return res
}
