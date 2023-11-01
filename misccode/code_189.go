package misccode

import "fmt"

func Rotate3(nums []int, k int) {
	pre := []int{}
	for i := len(nums) - k; i < len(nums); i++ {
		pre = append(pre, nums[i])
	}
	// copy(nums, append(pre, nums[:len(nums)-k]...))
	nums = append(pre, nums[:len(nums)-k]...)
	fmt.Println(nums)
}
func MapCopy(mm map[int]int) {
	// cm := map[int]int{2: 1}
	mm[2] = 1
}
