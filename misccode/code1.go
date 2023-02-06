package misccode

import "sort"

// 两数之和

// 1.暴力求解
func F1(nums []int, target int) [][]int {
	res := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		for j := i; j < len(nums); j++ {

			if nums[i] == nums[j] {
				res = append(res, []int{nums[i], nums[j]})
			}
		}
	}
	return res
}

// 2.增加存储降低时间复杂度
func F2(nums []int, target int) [][]int {

	res := make([][]int, 0)
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := m[target-nums[i]]; ok && m[target-nums[i]] != i {
			res = append(res, []int{nums[i], nums[v]})
		}
		m[nums[i]] = i
	}
	return res
}

func ThreeNumsSum(nums []int, target int) [][]int {

	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums); i++ {

		if nums[i] > target {
			break
		}
		// 去重
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		left, right := i+1, len(nums)-1
		for left < right {

			if nums[i]+nums[left]+nums[right] == target {
				res = append(res, []int{nums[i], nums[left], nums[right]})
				// 去重
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}
				left++
				right--
			} else if nums[i]+nums[left]+nums[right] > target {
				right--
			} else {
				left++
			}
		}
	}
	return res
}
