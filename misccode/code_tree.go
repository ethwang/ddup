package misccode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func dfs(root *TreeNode, prevSum int) int {
	if root == nil {
		return 0
	}
	sum := prevSum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		return sum
	}
	lsum := dfs(root.Left, sum)
	rsum := dfs(root.Right, sum)
	return lsum + rsum
}

func SumNumbers(root *TreeNode) int {
	return dfs(root, 0)
}
