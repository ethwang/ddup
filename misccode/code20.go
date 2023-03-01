package misccode

type MyStack struct {
	q1 *queen
	q2 *queen
}

func Constructor() MyStack {
	return MyStack{
		q1: newQueen(),
		q2: newQueen(),
	}

}

func (this *MyStack) Push(x int) {
	this.q2.push(x)
	for !this.q1.isEmpty() {
		this.q2.push(this.q1.pop())
	}
	for !this.q2.isEmpty() {
		this.q1.push(this.q2.pop())
	}
}

func (this *MyStack) Pop() int {

	return this.q1.pop()
}

func (this *MyStack) Top() int {

	return this.q1.top()
}

func (this *MyStack) Empty() bool {
	return this.q1.isEmpty()
}

type queen struct {
	data  []int
	front int
	tail  int
}

func newQueen() *queen {
	return &queen{
		data:  make([]int, 0),
		front: 0,
		tail:  0,
	}
}
func (q *queen) push(x int) {
	q.data = append(q.data, x)
	q.tail++
}
func (q *queen) size() int {
	return q.tail - q.front
}
func (q *queen) top() int {
	return q.data[q.front]
}
func (q *queen) pop() int {
	data := q.data[q.front]
	q.front++
	return data
}
func (q *queen) isEmpty() bool {
	if q.size() <= 0 {
		return true
	}
	return false
}
