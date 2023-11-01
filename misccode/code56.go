package misccode

import "sort"

func MergeIntervals(intervals [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	left, right := intervals[0][0], intervals[0][1]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] > right {
			ans = append(ans, []int{left, right})
			left, right = intervals[i][0], intervals[i][1]
		} else {
			right = max(intervals[i][1], right)
		}
	}
	ans = append(ans, []int{left, right})
	return ans
}
