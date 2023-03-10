package misccode

import "math"

func MinSubArrayLen(target int, nums []int) int {
	minLen := math.MaxInt
	win := nums[0]
	// 滑动窗口：如果窗口内值大于等于target缩小窗口(即左指针往右移),否则右指针往右移
	for i, j := 0, 0; i < len(nums) && j < len(nums); {
		if win < target {
			j++
			win += nums[j]
		} else {
			if j-i+1 < minLen {
				minLen = j - i + 1
			}
			win -= nums[i]
			i++
		}
	}
	if minLen == math.MaxInt {
		minLen = 0
	}
	return minLen
}
