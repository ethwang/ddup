package misccode

func Rotate(nums []int, k int) {
	for i := 0; i < k; i++ {
		tmp := []int{}
		tmp = append(tmp, nums[len(nums)-1])
		tmp = append(tmp, nums[0:len(nums)-1]...)
		copy(nums, tmp)
		//nums = tmp
	}
}
