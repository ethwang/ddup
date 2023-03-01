package misccode

func Trap(height []int) int {

	count := 0
	for i := 1; i < len(height)-1; i++ {

		lm, rm := height[i], height[i]
		for left := i - 1; left >= 0; left-- {
			if height[left] > lm {
				lm = height[left]
			}
		}
		for right := i + 1; right < len(height); right++ {
			if height[right] > rm {
				rm = height[right]
			}
		}
		if lm > height[i] && rm > height[i] {

			if lm < rm {
				count += lm - height[i]
			} else {
				count += rm - height[i]
			}
		}
	}
	return count
}
func Trap2(height []int) int {
	count := 0
	leftMax := make([]int, len(height))
	rightMax := make([]int, len(height))
	leftMax[0] = height[0]
	for i := 1; i < len(height); i++ {
		leftMax[i] = max(height[i-1], leftMax[i-1])
	}
	rightMax[len(height)-1] = height[len(height)-1]
	for j := len(height) - 2; j >= 0; j-- {
		rightMax[j] = max(rightMax[j-1], height[j-1])
	}
	for i := 1; i < len(height)-1; i++ {
		if leftMax[i] > height[i] && rightMax[i] > height[i] {
			count += min(leftMax[i], rightMax[i]) - height[i]
		}
	}
	return count
}

const MAX_INT = int(^uint(0) >> 1)
const MIN_INT = ^MAX_INT

func min(nums ...int) int {

	minNum := MAX_INT
	for _, v := range nums {
		if v < minNum {
			minNum = v
		}
	}
	return minNum
}
func max(nums ...int) int {
	maxNum := MIN_INT
	for _, v := range nums {
		if v > maxNum {
			maxNum = v
		}
	}
	return maxNum
}
