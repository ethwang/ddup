package treecode

import (
	"c1/queuecode"
	"c1/stackcode"
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// PreOrder 前序递归遍历
func PreOrder(head *TreeNode) {
	if head == nil {
		return
	}
	fmt.Println(head.Val)
	PreOrder(head.Left)
	PreOrder(head.Right)
}

// PreOrder2 前序非递归遍历
func PreOrder2(head *TreeNode) {
	sk := new(stackcode.Stack)
	if head != nil {
		sk.Push(head)
		for !sk.IsEmpty() {
			vTN, _ := sk.Pop().(TreeNode)
			fmt.Printf("%v ", vTN.Val)

			if vTN.Right != nil {
				sk.Push(vTN.Right)
			}
			if vTN.Left != nil {
				sk.Push(vTN.Left)
			}
		}
	}
}

// InOrder 中序递归遍历
func InOrder(head *TreeNode) {

	if head == nil {
		return
	}
	InOrder(head.Left)
	fmt.Println(head.Val)
	InOrder(head.Right)
}

// InOrder2 中序非递归遍历
func InOrder2(head *TreeNode) {
	sk := new(stackcode.Stack)
	h := head
	if h != nil {
		for !sk.IsEmpty() || h != nil {
			if h != nil {
				sk.Push(h)
				h = h.Left
			} else {
				vTN, _ := sk.Pop().(TreeNode)
				fmt.Printf("%v ", vTN.Val)
				h = vTN.Right
			}
		}
	}
}

// PostOrder 后序递归遍历
func PostOrder(head *TreeNode) {
	if head == nil {
		return
	}
	PostOrder(head.Left)
	PostOrder(head.Right)
	fmt.Println(head.Val)
}

// PostOrder2 后续非递归遍历
func PostOrder2(head *TreeNode) {
	h := head
	sk1 := new(stackcode.Stack)
	sk2 := new(stackcode.Stack)
	if h != nil {
		sk1.Push(h)
		for sk1 != nil {
			vTN, _ := sk1.Pop().(TreeNode)
			sk2.Push(vTN)
			if vTN.Left != nil {
				sk1.Push(vTN.Left)
			}
			if vTN.Right != nil {
				sk1.Push(vTN.Right)
			}
		}
	}
	for !sk2.IsEmpty() {
		vvTN, _ := sk2.Pop().(TreeNode)
		fmt.Printf("%v ", vvTN.Val)
	}
}

// 二叉树宽度优先搜索遍历
func BFS(head *TreeNode) {

	m := make(map[*TreeNode]int)
	q := new(queuecode.Queue)
	if head != nil {
		curLevel := 0
		curWidth := 0
		maxWidth := 0.0
		m[head] = 1
		q.Add(head)
		for !q.IsEmpty() {
			vTN, _ := q.Take().(*TreeNode)
			fmt.Printf("%v ", vTN.Val)
			if vTN.Left != nil {
				m[vTN.Left] = m[vTN] + 1
				q.Add(vTN.Left)
			}
			if vTN.Right != nil {
				m[vTN.Right] = m[vTN] + 1
				q.Add(vTN.Right)
			}
			if m[vTN] > curLevel {
				curWidth = 0
				curLevel = m[vTN]
			} else {
				curWidth++
			}
			maxWidth = math.Max(float64(maxWidth), float64(curWidth))
		}
		fmt.Printf("max width: %v", maxWidth)
	}
}

// 判断是否为搜索二叉树
func IsBST(head *TreeNode) bool {

	nums := make([]int, 0)
	isBST(head, nums)
	if len(nums) <= 0 {
		return false
	}
	pre := nums[0]
	for i := 1; i < len(nums); i++ {
		if pre > nums[i] {
			return false
		}
		pre = nums[i]
	}
	return true
}

func isBST(head *TreeNode, nums []int) {
	if head == nil {
		return
	}
	isBST(head.Left, nums)
	nums = append(nums, head.Val)
	isBST(head.Right, nums)
}

// IsCST是否是完全二叉树
func IsCST(head *TreeNode) bool {
	q := new(queuecode.Queue)
	if head != nil {
		flag := false
		q.Add(head)
		for !q.IsEmpty() {

			vTN, _ := q.Take().(*TreeNode)

			if (flag && (vTN.Left != nil || vTN.Right != nil)) || (vTN.Left == nil && vTN.Right != nil) {
				return false
			}
			if vTN.Left != nil {
				q.Add(vTN.Left)
			}
			if vTN.Right != nil {
				q.Add(vTN.Right)
			} else {
				flag = true
			}
		}
	}
	return true
}

// IsFullBT 是否是满二叉树
func IsFullBT(head *TreeNode) bool {

	if head == nil {
		return true
	}
	// 当前是第几层
	curLevel := 0
	// 当前节点在第几层
	nodeLevel := make(map[*TreeNode]int)
	// 当前层有多少个节点
	levelNodes := make(map[int]int)
	// 当前层的节点数
	widthNodes := 1
	q := new(queuecode.Queue)
	q.Add(head)
	nodeLevel[head] = 1
	for !q.IsEmpty() {
		trNode, _ := q.Take().(*TreeNode)
		if trNode.Left != nil {
			q.Add(trNode.Left)
			nodeLevel[trNode.Left] = nodeLevel[trNode] + 1

		}
		if trNode.Right != nil {
			q.Add(trNode.Right)
			nodeLevel[trNode.Right] = nodeLevel[trNode] + 1
		}

		if nodeLevel[trNode] > curLevel {
			curLevel = nodeLevel[trNode]
			widthNodes = 1
		} else {
			widthNodes++
		}
		levelNodes[curLevel] = widthNodes
	}

	for level, width := range levelNodes {
		if width != 2>>(level-1) {
			return false
		}
	}
	return true
}

// IsFullBT2 是否为满二叉树(递归套路)
func IsFullBT2(head *TreeNode) bool {
	isFull, _, _ := isFullBT(head)

	return isFull
}
func isFullBT(head *TreeNode) (bool, int, int) {

	if head == nil {
		return true, 0, 0
	}
	isFullLeft, countLeft, heightLeft := isFullBT(head.Left)
	isFullRight, countRight, heightRight := isFullBT(head.Left)
	maxHeight := math.Max(float64(heightLeft), float64(heightRight))

	return isFullLeft && isFullRight && (countLeft == countRight) && (heightLeft == heightRight), countLeft + countRight + 1, int(maxHeight) + 1
}

// 是否为平衡二叉树(递归套路)
func IsBalancedTree(head *TreeNode) bool {
	isBT, _ := isBalancedTree(head)
	return isBT
}

func isBalancedTree(head *TreeNode) (bool, int) {
	if head == nil {
		return true, 0
	}

	isBTLeft, hLeft := isBalancedTree(head.Left)
	isBtRight, hRight := isBalancedTree(head.Right)
	val := math.Abs(float64(hLeft) - float64(hRight))
	maxHeight := math.Max(float64(hLeft), float64(hRight))
	return isBTLeft && isBtRight && (val <= 1), int(maxHeight) + 1
}

// LowestCommonAncestor查找最低公共祖先
func LowestCommonAncestor1(head, n1, n2 *TreeNode) *TreeNode {

	fatherMap := make(map[*TreeNode]*TreeNode)
	getAncestor(head, fatherMap)

	fatherMap[head] = nil
	cur := n1
	setN1 := make(map[*TreeNode]interface{})
	for cur != nil {
		setN1[cur] = nil
		cur = fatherMap[cur]
	}

	cur = n2
	for cur != nil {
		if _, ok := setN1[cur]; ok {
			return cur
		}
		cur = fatherMap[cur]
	}
	return nil
}

func getAncestor(head *TreeNode, fatherM map[*TreeNode]*TreeNode) {

	if head == nil {
		return
	}
	fatherM[head.Left] = head
	fatherM[head.Right] = head
	getAncestor(head.Left, fatherM)
	getAncestor(head.Right, fatherM)
}

func LowestCommonAncestor2(head, n1, n2 *TreeNode) *TreeNode {
	if head == nil || head == n1 || head == n2 {
		return head
	}
	Left := LowestCommonAncestor2(head.Left, n1, n2)
	Right := LowestCommonAncestor2(head.Right, n1, n2)

	if Left != nil && Right != nil {
		return head
	}
	if Left != nil && Right == nil {
		return Left
	}
	if Left == nil && Right != nil {
		return Right
	}
	return nil
}

// 折纸(对折)问题
func Zhezhi(n int) {

	zhezhi(0, n, true)
}
func zhezhi(i, n int, data bool) {

	if i > n {
		return
	}
	zhezhi(i+1, n, true)
	fmt.Println(data)
	zhezhi(i+1, n, false)
}

// 创建完全二叉树
func CreateCST(n int) *TreeNode {

	root := preCreateCST(0, n, 0)
	return root
}
func preCreateCST(i, n, val int) *TreeNode {
	if i > n {
		return nil
	}
	node := &TreeNode{Val: val}
	node.Left = preCreateCST(i+1, n, 0)
	node.Right = preCreateCST(i+1, n, 1)

	return node
}

func PrintCst(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("%v,", root.Val)
	PrintCst(root.Left)
	PrintCst(root.Right)
}
