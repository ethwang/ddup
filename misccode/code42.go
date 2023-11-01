package misccode

func Trap42(height []int) int {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	ans := 0
	// 为啥是小于，而不是小于等于
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if leftMax < rightMax {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return ans
}
