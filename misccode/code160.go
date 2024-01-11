package misccode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	ha, hb := headA, headB
	haLen, hbLen := 0, 0
	for ha != nil {
		haLen++
		ha = ha.Next
	}
	for hb != nil {
		hbLen++
		hb = hb.Next
	}
	ha, hb = headA, headB
	if haLen-hbLen > 0 {
		ll := haLen - hbLen
		for ll > 0 {
			ha = ha.Next
			ll--
		}
	} else {
		ll := hbLen - haLen
		for ll > 0 {
			hb = hb.Next
			ll--
		}
	}

	for ha != nil && hb != nil {
		if ha == hb {
			return ha
		}
		ha, hb = ha.Next, hb.Next
	}
	return nil
}
func getIntersectionNode2(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	pa, pb := headA, headB
	for pa != pb {
		if pa == nil {
			pa = headB
		} else {
			pa = pa.Next
		}
		if pb == nil {
			pb = headA
		} else {
			pb = pb.Next
		}
	}
	return pa
}
