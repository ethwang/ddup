package linklistcode

type RandomNode struct {
	Val    int
	Random *RandomNode
	Next   *RandomNode
}

func InitRandomList(nums []int) (head *RandomNode) {
	head = &RandomNode{Val: nums[0]}
	preNode := head
	for j := 1; j < len(nums); j++ {
		node := &RandomNode{Val: nums[j]}
		preNode.Next = node
		preNode = node
	}

	h := head
	h.Random = h.Next.Next
	h.Next.Random = h

	return
}

// 深度拷贝有随机指针的链表(时间复杂度O(N),空间复杂度(O(N)))
func RandomNodeCopy(head *RandomNode) *RandomNode {

	m := make(map[*RandomNode]*RandomNode)
	h := head
	for h != nil {
		newNode := &RandomNode{Val: h.Val}
		m[h] = newNode
		h = h.Next
	}
	h1 := head
	for h1 != nil {
		m[h1].Next = m[h1.Next]
		m[h1].Random = m[h1.Random]
		h1 = h1.Next
	}
	return m[head]
}
func RandomNodeCopy2(head *RandomNode) (headCopy *RandomNode) {

	/*
		// 如果没有随机指针，下面链表的深度copy是可行的
		h := head
		if h != nil {
			headCopy = &RandomNode{Val: h.Val}
			preNode := headCopy
			for h2 := h.Next; h2 != nil; h2 = h2.Next {
				node := &RandomNode{Val: h2.Val}
				preNode.Next = node
				preNode = node
			}
		}
		 return
	*/

	// 利用原有链表节点相对位置关系，将新链表节点插入到原有链表中
	h := head
	for h != nil {

		next := h.Next
		newNode := &RandomNode{Val: h.Val}
		h.Next = newNode
		newNode.Next = next
		h = next
	}
	h = head
	for h != nil {
		next := h.Next.Next
		h.Next.Random = h.Random.Next
		h = next
	}
	h, headCopy = head.Next, head.Next

	for h != nil {
		h.Next = h.Next.Next
		h = h.Next.Next
	}
	return
}
