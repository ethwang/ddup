package misccode

func Subsets(nums []int) [][]int {
	result := [][]int{}
	path := make([]int, 0, 4)
	backtrackingslice(0, nums, &path, &result)
	return result
}
func backtrackingslice(startIndex int, nums []int, path *[]int, result *[][]int) {
	*result = append(*result, *path)
	if startIndex == len(nums) {
		return
	}
	for i := startIndex; i < len(nums); i++ {
		*path = append(*path, nums[i])
		backtrackingslice(i+1, nums, path, result)
		*path = (*path)[:len(*path)-1]
	}
}
