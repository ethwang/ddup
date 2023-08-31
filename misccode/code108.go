package misccode

func SortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	val := nums[len(nums)/2]
	head := &TreeNode{Val: val}

	head.Left = SortedArrayToBST(nums[:len(nums)/2])
	if len(nums)/2+1 >= len(nums) {
		return nil
	}
	head.Right = SortedArrayToBST(nums[len(nums)/2+1:])
	return head
}
