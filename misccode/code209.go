package misccode

import (
	"math"
	"sort"
)

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

func MinSubArrayLen2(target int, nums []int) int {
	presum := map[int]int{0: len(nums)}
	sum := 0
	minLen := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for j := 0; j <= sum-target; j++ {

			if presum[sum-(target+j)] > 0 {
				minLen = min(minLen, i-presum[sum-(target+j)])
			}
		}
		presum[sum] = i
	}
	if minLen == math.MaxInt {
		return 0
	}
	return minLen
}

func MinSubArrayLen3(s int, nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	ans := math.MaxInt32
	sums := make([]int, n+1)
	// 为了方便计算，令 size = n + 1
	// sums[0] = 0 意味着前 0 个元素的前缀和为 0
	// sums[1] = A[0] 前 1 个元素的前缀和为 A[0]
	// 以此类推
	for i := 1; i <= n; i++ {
		sums[i] = sums[i-1] + nums[i-1]
	}
	for i := 1; i <= n; i++ {
		target := s + sums[i-1]
		bound := sort.SearchInts(sums, target)
		if bound < 0 {
			bound = -bound - 1
		}
		if bound <= n {
			ans = min(ans, bound-(i-1))
		}
	}
	if ans == math.MaxInt32 {
		return 0
	}
	return ans
}
