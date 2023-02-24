package misccode

func LengthOfLIS(nums []int) int {

	path := make([]int, 0)
	length := 0
	backtracking18(0, nums, path, &length)
	return length
}
func backtracking18(startIndex int, nums []int, path []int, length *int) {
	if len(path) > *length {
		*length = len(path)
	}
	if startIndex == len(nums) {
		return
	}
	for i := startIndex; i < len(nums); i++ {
		if len(path) == 0 || nums[i] > path[len(path)-1] {
			path = append(path, nums[i])
			backtracking18(i+1, nums, path, length)
			path = path[:len(path)-1]
		}
	}
}
