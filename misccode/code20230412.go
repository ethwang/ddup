package misccode

// 1,2,3,5,4,2
func SearchX(nums []int) int {
	left := 0
	right := len(nums) - 1
	for left < right {
		mid := (left + right) / 2
		if nums[mid-1] < nums[mid] {
			left = mid
		} else if nums[mid-1] > nums[mid] {
			right = mid - 1
		}
	}
	return left
}
