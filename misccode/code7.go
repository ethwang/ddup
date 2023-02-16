package misccode

type stack struct {
	arr    []byte
	top    int
	bottom int
}

func newStack() *stack {
	return &stack{
		arr:    make([]byte, 0),
		top:    0,
		bottom: 0,
	}
}
func (s *stack) isEmpty() bool {

	return s.top == 0
}

func (s *stack) pop() byte {
	var v byte
	if s.top > s.bottom {
		s.top--
		v = s.arr[s.top]
		s.arr = s.arr[:len(s.arr)-1]
	}
	return v
}
func (s *stack) push(ch byte) {
	s.arr = append(s.arr, ch)
	s.top++
}

func IsValid(s string) bool {

	m := map[byte]byte{
		'[': ']',
		'{': '}',
		'(': ')',
	}
	st := newStack()
	for i := 0; i < len(s); i++ {

		if _, ok := m[s[i]]; ok {
			st.push(s[i])
		} else {

			if v, ok := m[st.pop()]; !ok || v != s[i] {
				return false
			}
		}
	}
	return st.isEmpty()
}
