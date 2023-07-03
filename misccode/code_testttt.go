package misccode

import "math"

type ListNode1 struct {
	Val  int
	Next *ListNode1
}

func GetIntersectionNode(headA *ListNode1, headB *ListNode1) *ListNode1 {
	ha, hb := headA, headB
	counta, countb := 0, 0
	for ha != nil {
		ha = ha.Next
		counta++
	}
	for hb != nil {
		hb = hb.Next
		countb++
	}
	ha1, hb1 := headA, headB
	step := math.Abs(float64(counta) - float64(countb))
	if counta > countb {
		for step > 0 {
			ha1 = ha1.Next
			step--
		}
	} else {
		for step > 0 {
			hb1 = hb1.Next
			step--
		}
	}
	for ha1 != nil && hb1 != nil {
		if ha1 == hb1 {
			return ha1
		}
		ha1 = ha1.Next
		hb1 = hb1.Next
	}
	return nil
}
