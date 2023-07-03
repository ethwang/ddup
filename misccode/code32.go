package misccode

func LongestValidParentheses(s string) int {
	maxLen := 0
	stack := []int{-1, 0}
	for i := 1; i < len(s); i++ {
		if s[i] == ')' {
			if len(stack)-1 >= 0 && stack[len(stack)-1] >= 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
				if i-stack[len(stack)-1] > maxLen {
					maxLen = i - stack[len(stack)-1]
				}
			} else {
				stack = append(stack, i)
			}
		} else {
			stack = append(stack, i)
		}
	}
	return maxLen
}
