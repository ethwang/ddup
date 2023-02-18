package misccode

import "sort"

func LongestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sort.Ints(nums)
	length := 0
	longestLen := 0
	for i, j := 0, 1; i < len(nums) && j < len(nums); i, j = i+1, j+1 {
		if nums[i] == nums[j] {
			continue
		}
		if nums[i]+1 == nums[j] {
			length++
		} else {
			length = 0
		}
		if length > longestLen {
			longestLen = length
		}

	}
	return longestLen
}
func LongestConsecutive2(nums []int) int {
	m := make(map[int]bool)
	for _, v := range nums {
		m[v] = true
	}
	longestLen := 0
	for _, v := range nums {

		if m[v-1] {
			continue
		}
		curNum := v
		curLen := 1
		for m[curNum+1] {
			curNum++
			curLen++
		}
		if curLen > longestLen {
			longestLen = curLen
		}
	}
	return longestLen
}
