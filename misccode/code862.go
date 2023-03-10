package misccode

import "math"

func ShortestSubarray(nums []int, k int) int {
	// map[k]v k是前i个数的和,v是该k下的最长数组下标

	presum := map[int]int{0: -1}
	sum := 0
	minlen := math.MaxInt
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		for ; sum-k >= 0; k++ {
			if v, ok := presum[sum-k]; ok {
				minlen = min(i-v, minlen)
			}
		}

		presum[sum] = i
	}
	return minlen
}
