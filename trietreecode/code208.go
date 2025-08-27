package trietreecode

type Trie struct {
	Node [26]*Trie
}

func Constructor() Trie {
	return Trie{
		Node: [26]*Trie{},
	}

}

func (this *Trie) Insert(word string) {
	//curNode := this.Node
	//for i := 0; i < len(word); i++ {
	//	if curNode[word[i]-'a'] == nil {
	//		node := Constructor()
	//		curNode[word[i]-'a'] = &node
	//	}
	//	curNode = curNode[word[i]-'a'].Node
	//}
	curNode := this
	for _, ch := range word {
		ch -= 'a'
		if curNode.Node[ch] == nil {
			next := Constructor()
			curNode.Node[ch] = &next
		}
		curNode = curNode.Node[ch]
	}
}

func (this *Trie) Search(word string) bool {
	curNode := this.Node
	for i := 0; i < len(word); i++ {
		if curNode[word[i]-'a'] == nil {
			return false
		}
		curNode = curNode[word[i]-'a'].Node
	}
	return true
}

func (this *Trie) StartsWith(prefix string) bool {
	curNode := this.Node
	for i := 0; i < len(prefix); i++ {
		if curNode[prefix[i]-'a'] != nil {
			curNode = curNode[prefix[i]-'a'].Node
		} else {
			return false
		}
	}
	return true
}
