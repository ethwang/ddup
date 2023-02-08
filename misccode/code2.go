package misccode

import (
	"math/rand"
	"time"
)

func QuickSort2(nums []int) {
	quickSort2(nums, 0, len(nums)-1)
}

func quickSort2(nums []int, i, j int) {
	if i >= j {
		return
	}
	// 随机种子初始化random
	rd := rand.New(rand.NewSource(time.Now().UnixNano()))
	x := i + rd.Intn(j-i+1)
	nums[x], nums[j] = nums[j], nums[x]
	l, r := t1(nums, i, j)
	quickSort2(nums, i, l)
	quickSort2(nums, r+1, j)
}
func t1(nums []int, i, j int) (int, int) {
	left := i - 1
	right := j
	for i < right {
		if nums[i] > nums[j] {
			right--
			exchange(nums, i, right)
		} else if nums[i] < nums[j] {
			left++
			exchange(nums, i, left)
			i++
		} else {
			i++
		}
	}
	exchange(nums, j, right)
	return left, right
}

func exchange(nums []int, i, j int) {
	tmp := nums[i]
	nums[i] = nums[j]
	nums[j] = tmp
}
