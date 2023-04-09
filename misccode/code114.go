package misccode

func Flatten(root *TreeNode) {
	ans := []*TreeNode{}
	inorder(root, &ans)
	for i := 0; i < len(ans)-1; i++ {
		ans[i].Right = ans[i+1]
	}
	root = ans[0]
}
func inorder(root *TreeNode, ans *[]*TreeNode) {
	if root == nil {
		return
	}
	newNode := &TreeNode{Val: root.Val}
	*ans = append(*ans, newNode)
	inorder(root.Left, ans)
	inorder(root.Right, ans)
}
