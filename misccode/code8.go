package misccode

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	head := &ListNode{}
	h := head

	for list1 != nil && list2 != nil {

		if list1.Val < list2.Val {
			h.Next = list1
			h = list1
			list1 = list1.Next
		} else {
			h.Next = list2
			h = list2
			list2 = list2.Next
		}
	}
	if list1 != nil {
		h.Next = list1
	}
	if list2 != nil {
		h.Next = list2
	}

	return head.Next
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	return divide(lists, 0, len(lists)-1)
}
func divide(lists []*ListNode, left, right int) *ListNode {
	if left == right {
		return lists[left]
	}
	mid := (left + right) / 2
	l := divide(lists, left, mid)
	r := divide(lists, mid+1, right)
	return mergeTwoLists(l, r)
}
