package misccode

import "sort"

func SubsetsWithDup(nums []int) [][]int {

	sort.Ints(nums)
	r := make([]int, 0)
	res := make([][]int, 0)
	process15(false, 0, nums, r, &res)
	return res
}

func process15(choosePre bool, start int, nums []int, r []int, res *[][]int) {

	//dst := make([]int, len(r))
	//copy(dst, r)
	//*res = append(*res, dst)
	////if start == len(nums) {
	////	dst := make([]int, len(r))
	////	copy(dst, r)
	////	*res = append(*res, dst)
	////	return
	////}
	//for i := start; i < len(nums); i++ {

	//	r = append(r, nums[i])
	//	process15(i+1, nums, r, res)
	//	r = r[:len(r)-1]
	//}
	if start == len(nums) {

		dst := make([]int, len(r))
		copy(dst, r)
		*res = append(*res, dst)
		return
	}

	if !choosePre && start > 0 && nums[start-1] == nums[start] {
		return
	}
	r = append(r, nums[start])
	process15(true, start+1, nums, r, res)
	r = r[:len(r)-1]
	process15(false, start+1, nums, r, res)
}
