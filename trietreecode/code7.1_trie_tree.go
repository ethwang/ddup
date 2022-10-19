package trietreecode

/*
	前缀树结构
*/
type TrieTreeNode struct {
	Path  int
	End   int
	Nodes []*TrieTreeNode
}

/*
	前缀树 功能
	新增
	删除
	返回字符串前缀数量
*/

var TrieTree = &TrieTreeNode{
	Nodes: make([]*TrieTreeNode, 26),
}

func Insert(s string) {

	bs := []byte(s)
	node := TrieTree
	node.Path++
	for _, v := range bs {
		d := v - 'a'
		if node.Nodes[d] == nil {
			node.Nodes[d] = &TrieTreeNode{
				Nodes: make([]*TrieTreeNode, 26),
			}
		}
		node = node.Nodes[d]
		node.Path++
	}
	node.End++

	// fmt.Println(TrieTree)
}

func Del(s string) {
	if isExist(s) > 0 {
		node := TrieTree
		node.Path--

		bs := []byte(s)
		for _, v := range bs {
			d := v - 'a'
			node = node.Nodes[d]
			node.Path--
		}
		node.End--
	}
}

func isExist(s string) int {
	bs := []byte(s)
	node := TrieTree
	for _, v := range bs {

		d := v - 'a'
		if node.Nodes[d] == nil {
			return 0
		}
		node = node.Nodes[d]
	}
	return node.End
}

func PreCount(pre string) int {
	bs := []byte(pre)

	node := TrieTree
	for _, v := range bs {

		d := v - 'a'
		if node.Nodes[d] != nil {
			node = node.Nodes[d]
		} else {
			return 0
		}
	}
	return node.Path
}
