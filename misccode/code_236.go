package misccode

func llowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var dfs func(*TreeNode)
	parent := map[*TreeNode]*TreeNode{}
	visited := map[*TreeNode]bool{}
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		if node.Left != nil {
			parent[node.Left] = node
			dfs(node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			dfs(node.Right)
		}
	}
	dfs(root)
	for p != nil {
		visited[parent[p]] = true
		p = parent[p]
	}
	for q != nil {
		if v, ok := visited[parent[q]]; ok && v {
			return parent[q]
		}
		q = parent[q]
	}
	return nil
}
