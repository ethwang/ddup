package queuecode

type Node struct {
	data interface{}
	next *Node
}

type Queue struct {
	head *Node
	tail *Node
	len  int
}

func (q *Queue) IsEmpty() bool {

	return q.len <= 0
}

func (q *Queue) Add(data interface{}) {
	newNode := &Node{data: data}

	if q.len == 0 {

		q.head = newNode
		q.tail = newNode
		q.len++
	} else {

		q.tail.next = newNode
		q.tail = newNode
		q.len++
	}
}
func (q *Queue) Take() interface{} {

	if q.len == 0 {
		return nil
	}
	d := q.head.data
	q.head = q.head.next
	q.len--

	return d
}
func (q *Queue) Peek() interface{} {

	if q.len == 0 {
		return nil
	}
	return q.head.data
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type Queen2 struct {
	q    []*TreeNode
	head int
	tail int
}

func NewQueen2() *Queen2 {
	return &Queen2{
		q:    make([]*TreeNode, 0),
		head: 0,
		tail: 0,
	}
}

func (q2 *Queen2) Size() int {
	return q2.tail - q2.head + 1
}

func (q2 *Queen2) IsEmpty() bool {
	return q2.tail < q2.head
}

func (q2 *Queen2) Push(tNode *TreeNode) {
	q2.q = append(q2.q, tNode)
	q2.tail = len(q2.q) - 1
}

func (q2 *Queen2) Pop() *TreeNode {
	if q2.tail < q2.head {
		return nil
	}
	val := q2.q[q2.head]
	q2.head = q2.head + 1
	return val
}
func LevelOrder2(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	q := NewQueen2()
	q.Push(root)
	res := make([][]int, 0)
	for !q.IsEmpty() {
		n := q.Size()
		tmp := make([]int, 0)
		for n > 0 {
			node := q.Pop()
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				q.Push(node.Left)
			}
			if node.Right != nil {
				q.Push(node.Right)
			}
			n--
		}
		res = append(res, tmp)
	}
	return res
}

func LevelOrder(root *TreeNode) [][]int {

	res := make([][]int, 0)
	if root == nil {
		return res
	}
	tns := make([]*TreeNode, 0)
	tns = append(tns, root)
	for len(tns) > 0 {
		n := len(tns)
		tmp := make([]int, 0)
		for i := 0; i < n; i++ {
			tmp = append(tmp, tns[i].Val)
			if tns[i].Left != nil {
				tns = append(tns, tns[i].Left)
			}
			if tns[i].Right != nil {
				tns = append(tns, tns[i].Right)
			}
		}
		res = append(res, tmp)
		tns = tns[n:]
	}
	return res
}
