package stackcode

type Stack struct {
	nums []interface{}
	len  int
}

func (sk *Stack) Pop() interface{} {
	if sk.len <= 0 {
		return nil
	}
	sk.len = sk.len - 1
	return sk.nums[sk.len]
}

func (sk *Stack) Push(data interface{}) {
	sk.nums = append(sk.nums, data)
	sk.len++
}

func (sk *Stack) Len() int {
	return sk.len
}

func (sk *Stack) IsEmpty() bool {
	return sk.Len() <= 0
}
