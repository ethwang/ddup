package misccode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	return nil
}
func preorder10(root, target *TreeNode, list *[]*TreeNode) {
	if root == target || root == nil {
		return
	}
	*list = append(*list, root)
	preorder10(root.Left, target, list)
	*list = (*list)[:len(*list)-1]
	preorder10(root.Right, target, list)
	*list = (*list)[:len(*list)-1]
}
