package misccode

import "sort"

func Merge(intervals [][]int) [][]int {
	// 1.按照每个区间首字符从小到大排序
	// 2.合并后塞入新数组并把原数组后续也一并塞入新数组将新数组赋值给原数组,并保持双指针不变
	// 3.如果没有合并，双指针分别向后移动一位

	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] == intervals[j][0] {
			return intervals[i][1] < intervals[j][1]
		}
		return intervals[i][0] < intervals[j][0]
	})
	res := make([][]int, 0)
	for i := 0; i < len(intervals); {
		j := i + 1

		if j >= len(intervals) {
			res = append(res, intervals[i])
			break
		}
		if intervals[j][0] <= intervals[i][1] {
			newIntervals := make([][]int, 0)
			newNums := make([]int, 0)
			if intervals[j][1] <= intervals[i][1] {
				newNums = append(newNums, intervals[i][0], intervals[i][1])
			} else {
				newNums = append(newNums, intervals[i][0], intervals[j][1])
			}
			newIntervals = append(newIntervals, newNums)
			for k := j + 1; k < len(intervals); k++ {
				newIntervals = append(newIntervals, intervals[k])
			}
			intervals = newIntervals
			i = 0
			continue
		}
		if intervals[j][0] > intervals[i][1] {
			res = append(res, intervals[i])
			i++
		}
	}
	return res
}
