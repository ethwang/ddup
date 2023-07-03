package misccode

func MaximalRectangle(matrix [][]byte) int {
	maxR := 0
	mat := make([][]int, len(matrix))
	for i := 0; i < len(mat); i++ {
		mat[i] = make([]int, len(matrix[i]))
	}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if i == 0 {
				mat[i][j] = int(matrix[i][j] - '0')
			} else if matrix[i][j] == '1' {
				mat[i][j] = mat[i-1][j] + 1
			}
		}
	}
	for i := 0; i < len(mat); i++ {
		maxR = max(maxR, maxRectangle(mat[i]))
	}
	return maxR
}
func maxRectangle(rh []int) int {
	maxArea := 0
	rh = append(rh, 0)
	stack := [][2]int{{-1, 0}} // 0:坐标,1:高度
	for i := 0; i < len(rh); i++ {
		for rh[i] < stack[len(stack)-1][1] {
			cur := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			maxArea = max(maxArea, cur[1]*(i-stack[len(stack)-1][0]-1))
		}
		stack = append(stack, [2]int{i, rh[i]})
	}
	return maxArea
}
