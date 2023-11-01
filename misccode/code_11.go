package misccode

func maxArea(height []int) int {
	// area=(j-i)*min(heigh[i],height[j])
	m := 0
	for i, j := 0, len(height)-1; i != j; i, j = i+1, j-1 {
		if m < (j-i)*min(height[i], height[j]) {
			m = (j - i) * min(height[i], height[j])
		}
		if height[i] > height[j] {
			j--
		} else {
			i++
		}
	}
	return m
}
