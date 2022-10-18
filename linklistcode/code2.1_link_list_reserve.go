package linklistcode

type Node struct {
	Value int
	Next  *Node
	Pre   *Node
}

// 初始化单向链表
func InitLinkList(nums []int) *Node {
	head := &Node{}
	for _, v := range nums {
		newNode := &Node{Value: v}

		nextNode := head.Next
		head.Next = newNode
		newNode.Next = nextNode
	}
	return head.Next
}

// 反转单向链表
func ReserveLinkList(h *Node) *Node {
	var preNode *Node
	head := h
	for head != nil {
		nextNode := head.Next
		head.Next = preNode
		preNode = head
		head = nextNode
	}
	return preNode
}

// 初始化双向链表
func InitTwoWayLinkList(nums []int) *Node {
	head := &Node{}
	for _, v := range nums {
		newNode := &Node{Value: v}

		nextNode := head.Next
		head.Next = newNode
		newNode.Next = nextNode

		if nextNode != nil {
			nextNode.Pre = newNode
		}
		newNode.Pre = head
	}
	return head
}
func InitTwoWayLinkListHasHeadAndTail(nums []int) (*Node, *Node) {
	head, tail := &Node{}, &Node{}
	head.Next = tail
	tail.Pre = head
	for _, v := range nums {
		newNode := &Node{Value: v}

		nextNode := head.Next
		head.Next = newNode
		newNode.Next = nextNode

		if nextNode != nil {
			nextNode.Pre = newNode
		}
		newNode.Pre = head
	}
	return head, tail
}

// 反转双向链表
func ReserveTwoWayLinkList(h *Node) *Node {
	var preNode *Node
	head := h.Next
	for head != nil {
		nextNode := head.Next
		head.Next = preNode
		head.Pre = nextNode
		preNode = head
		head = nextNode
	}
	return preNode
}
