package misccode

func jump(nums []int) int {
	if len(nums) <= 1 {
		return 0
	}
	step := 1
	for i := 0; i < len(nums); {
		if i+nums[i] >= len(nums)-1 {
			return step
		}
		maxIndex := i
		for j := i; j <= i+nums[i]; j++ {
			if nums[j]+j > nums[maxIndex]+maxIndex {
				maxIndex = j
			}
		}
		i = maxIndex
		step++
	}
	return step
}
