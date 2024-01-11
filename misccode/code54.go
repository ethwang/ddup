package misccode

func SpiralOrder(matrix [][]int) []int {
	ans := []int{}
	rLen := len(matrix)
	cLen := len(matrix[0])
	top, left, bottom, right := 0, 0, rLen-1, cLen-1
	for top <= bottom && left <= right {
		for i := left; i <= right; i++ {
			ans = append(ans, matrix[top][i])
		}
		for i := top + 1; i <= bottom; i++ {
			ans = append(ans, matrix[i][right])
		}
		if top < bottom && left < right {
			for i := right - 1; i >= left; i-- {
				ans = append(ans, matrix[bottom][i])
			}
			for i := bottom - 1; i > left; i-- {
				ans = append(ans, matrix[i][left])
			}
		}

		top++
		left++
		bottom--
		right--
	}
	return ans
}
