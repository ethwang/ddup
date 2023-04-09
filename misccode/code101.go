package misccode

func IsSymmetric(root *TreeNode) bool {
	q := []*TreeNode{}
	q = append(q, root)
	level := 0
	for len(q) > 0 {
		n := len(q)
		if !isSymme(q, level) {
			return false
		}
		for i := 0; i < n; i++ {
			if q[i].Left != nil {
				q = append(q, q[i].Left)
			}
			if q[i].Right != nil {
				q = append(q, q[i].Right)
			}
		}
		level++
		q = q[n:]
	}
	return true
}
func isSymme(nodes []*TreeNode, level int) bool {
	if 2<<level != len(nodes) {
		return false
	}
	for i, j := 0, len(nodes)-1; i < j; i, j = i+1, j-1 {
		if nodes[i].Val != nodes[j].Val {
			return false
		}
	}
	return true
}
