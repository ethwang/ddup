package linklistcode

func IsPalindrome(head *Node) bool {
	len := getLinkListLength(head)

	newHead, nextNode := reserveLinkListWithLen(head, len/2)
	if len%2 == 0 {

		for newHead != nil && nextNode != nil {
			if newHead.Value != nextNode.Value {
				return false
			}
			newHead = newHead.Next
			nextNode = nextNode.Next
		}
	} else {
		if nextNode != nil {
			newHead2 := nextNode.Next
			for newHead != nil && newHead2 != nil {
				if newHead.Value != newHead2.Value {
					return false
				}
				newHead = newHead.Next
				newHead2 = newHead2.Next
			}
		}
	}
	return true
}
func reserveLinkListWithLen(head *Node, len int) (newHead, nextNode *Node) {
	h := head
	i := 0
	for h != nil && i < len {
		nextNode = h.Next
		h.Next = newHead
		newHead = h
		h = nextNode
		i++
	}
	return
}
func getLinkListLength(head *Node) (len int) {
	h := head
	for h != nil {
		len++
		h = h.Next
	}
	return
}
