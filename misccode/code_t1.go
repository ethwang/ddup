package misccode

import "fmt"

// [1,2,3,4,5] 6

// 1st: left=0 right=4 mid = 2  nums[mid]=3
// 2nd: left=3 right=4 mid = 3 nums[mid] = 4
// 3rd: left=4 right=4 mid=4 nums[mid]=5
// 4 :  left=5 right=4 mid=4 nums[mid]=5
func SearchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := (left + right) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

// [1,3,4,5,6] 2
func Insert() {

	nums := []int{1, 3, 4, 5, 6}
	target := 7
	index := SearchInsert(nums, target)

	newNums := make([]int, 0)

	// 5
	pre := nums[:index]
	after := nums[index:]

	newNums = append(newNums, pre...)
	newNums = append(newNums, target)
	newNums = append(newNums, after...)
	fmt.Println(newNums)
}
