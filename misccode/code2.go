package misccode

func t1(nums []int, i, j int) {
	left := i - 1
	right := j
	k := nums[j]
	for i := 0; i < len(nums); {

		for left < right {
			if nums[i] > nums[k] {
				right--
				exchange(nums, i, right)
			} else {
				left++
				i++
			}
		}
	}
}
func exchange(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
