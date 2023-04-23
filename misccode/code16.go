package misccode

import (
	"math"
	"sort"
)

func ThreeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	sum := math.MaxInt

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r]+nums[i] == target {
				return target
			} else if nums[l]+nums[r]+nums[i] > target {
				sum = min(nums[l]+nums[r]+nums[i]-target, sum)

				r0 := r - 1
				for l < r0 && nums[r0] == nums[r] {
					r0--
				}
				r = r0
			} else {
				sum = min(target-(nums[l]+nums[r]+nums[i]), sum)

				l0 := l + 1
				for l0 < r && nums[l0] == nums[l] {
					l0++
				}
				l = l0
			}
		}
	}
	return sum
}
