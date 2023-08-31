package misccode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func ReverseKGroup(head *ListNode, k int) *ListNode {
	tmp := []*ListNode{}
	th := head
	for th != nil {
		tmp = append(tmp, th)
		th = th.Next
	}
	for _, node := range tmp {
		node.Next = nil
	}

	for i := range tmp {
		if (i+1)%k == 0 {
			exchange2(tmp[i+1-k : i+1])
		}
	}
	dummyHead := &ListNode{}
	th = dummyHead
	for _, node := range tmp {
		th.Next = node
		th = th.Next
	}
	return dummyHead.Next
}
func exchange2(arr []*ListNode) {
	for i := 0; i < len(arr)/2; i++ {
		arr[i], arr[len(arr)-1-i] = arr[len(arr)-1-i], arr[i]
	}
}
