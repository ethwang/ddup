package misccode

func Exist(board [][]byte, word string) bool {

	direction := [][]int{{-1, 0}, {0, -1}, {1, 0}, {0, 1}}
	used := make([][]int, len(board))
	for i := 0; i < len(board); i++ {
		used[i] = make([]int, len(board[i]))
	}
	for m := 0; m < len(board); m++ {
		for n := 0; n < len(board[m]); n++ {
			if len(word) == 1 && board[m][n] == word[0] {
				return true
			}
			if backtracking(0, m, n, board, word, direction, used) {
				return true
			}
		}
	}
	return false
}
func backtracking(index, i, j int, board [][]byte, word string, direction [][]int, used [][]int) bool {

	used[i][j] = 1
	defer func() {
		used[i][j] = 0
	}()
	if word[index] != board[i][j] {
		return false
	}
	if index == len(word)-1 {
		return true
	}
	for d := 0; d < len(direction); d++ {
		newI := i + direction[d][0]
		newJ := j + direction[d][1]
		if newI >= 0 && newI <= len(board)-1 && newJ >= 0 && newJ <= len(board[i])-1 {
			if used[newI][newJ] == 1 {
				continue
			}
			if backtracking(index+1, newI, newJ, board, word, direction, used) {
				return true
			}
		}
	}
	return false
}
