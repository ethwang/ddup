package misccode

func SearchMatrix(matrix [][]int, target int) bool {
	// 二分
	for i := 0; i < len(matrix); i++ {
		l, r := 0, len(matrix[i])-1
		for l <= r {
			m := (l + r) / 2
			if matrix[i][m] == target {
				return true
			} else if matrix[i][m] > target {
				r = m - 1

			} else {
				l = m + 1
			}
		}
	}
	return false
}
