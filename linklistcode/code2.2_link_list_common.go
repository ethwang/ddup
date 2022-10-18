package linklistcode

func retCircleNode(head *Node) (circleNode *Node) {

	h := head
	m := make(map[*Node]bool)

	for h != nil {
		if !m[h] {
			m[h] = true
		} else {
			circleNode = h
		}
		h = h.Next
	}
	return
}

func RetTwoLinkListCommonHeadWithCircle(head1, head2 *Node) (commonHead *Node) {
	h1, h2 := head1, head2
	circleNode1 := retCircleNode(h1)
	circleNode2 := retCircleNode(h2)
	if circleNode1 == circleNode2 {
		h1, h2 = head1, head2
		h1Len, h2Len := 0, 0
		for h1 != circleNode1 {
			h1Len++
			h1 = h1.Next
		}
		for h2 != circleNode1 {
			h2Len++
			h2 = h2.Next
		}
		h1, h2 = head1, head2
		diffLen := 0
		if h1Len > h2Len {
			diffLen = h1Len - h2Len
			for diffLen > 0 {
				h1 = h1.Next
				diffLen--
			}
		} else {
			diffLen = h2Len - h1Len
			for diffLen > 0 {
				h2 = h2.Next
				diffLen--
			}
		}
		for h1 != circleNode1 && h2 != circleNode1 {
			if h1 == h2 {
				commonHead = h1
				break
			}
			h1 = h1.Next
			h2 = h2.Next
		}
		return
	} else {
		c1 := circleNode1.Next
		for c1 != circleNode1 {
			if c1 == circleNode2 {
				commonHead = c1
				return
			}
			c1 = c1.Next
		}
		return
	}
	//if circleNode1 != nil && circleNode2 == nil {
	//	h1, h2 = head1, head2

	//	h1Len, h2Len := 0, 0
	//	for h1 != circleNode1 {
	//		h1Len++
	//		h1 = h1.Next
	//	}
	//	for h2 != circleNode1 {
	//		h2Len++
	//		h2 = h2.Next
	//	}
	//	diffLen := 0
	//	h1New, h2New := head1, head2
	//	if h2Len > h1Len {
	//		diffLen = h2Len - h1Len
	//		for diffLen > 0 && h2New != circleNode1 {
	//			h2New = h2New.Next
	//			diffLen--
	//		}
	//	} else {
	//		diffLen = h1Len - h2Len
	//		for diffLen > 0 && h1New != circleNode1 {
	//			h1New = h1New.Next
	//			diffLen--
	//		}
	//	}

	//	for h1New != nil && h2New != nil {

	//		if h1New.Value == h2New.Value {
	//			commonHead = h1New
	//			break
	//		}
	//		h1New = h1New.Next
	//		h2New = h2New.Next
	//	}

	//}
}

func RetTwoLinkListCommonHeadWithoutCircle(head1, head2 *Node) (commonHead *Node) {
	h1, h2 := head1, head2
	h1Len, h2Len := 0, 0
	for h1 != nil {
		h1Len++
		h1 = h1.Next
	}
	for h2 != nil {
		h2Len++
		h2 = h2.Next
	}
	diffLen := 0
	h1New, h2New := head1, head2
	if h2Len > h1Len {
		diffLen = h2Len - h1Len
		for diffLen > 0 && h2New != nil {
			h2New = h2New.Next
			diffLen--
		}
	} else {
		diffLen = h1Len - h2Len
		for diffLen > 0 && h1New != nil {
			h1New = h1New.Next
			diffLen--
		}
	}

	for h1New != nil && h2New != nil {

		if h1New == h2New {
			commonHead = h1New
			break
		}
		h1New = h1New.Next
		h2New = h2New.Next
	}

	return
}
