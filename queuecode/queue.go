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
