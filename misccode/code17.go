package misccode

func NumIslands(grid [][]byte) int {
	count := 0
	used := make([][]int, len(grid))
	for i := 0; i < len(grid); i++ {
		used[i] = make([]int, len(grid[i]))
	}
	direction := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == '1' && used[i][j] == 0 {
				backtracking17(i, j, grid, direction, &used)
				count++
			}
		}
	}
	return count
}
func backtracking17(row, column int, grid [][]byte, direction [][]int, used *[][]int) {

	for i := 0; i < len(direction); i++ {
		newI := row + direction[i][0]
		newJ := column + direction[i][1]
		if newI >= 0 && newI < len(grid) && newJ >= 0 && newJ < len(grid[0]) && grid[newI][newJ] == '1' && (*used)[newI][newJ] == 0 {
			(*used)[newI][newJ] = 1
			backtracking17(newI, newJ, grid, direction, used)
		}
	}
}
