package misccode

func DecodeString(s string) string {
	stack := make([]byte, 0)
	for i := 0; i < len(s); i++ {
		if s[i] != ']' {
			stack = append(stack, s[i])
		} else {

			stmp := ""
			for len(stack) > 0 && stack[len(stack)-1] != '[' {
				stmp = string(stack[len(stack)-1]) + stmp
				stack = stack[:len(stack)-1]
			}
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			}
			count := 0
			i := 1
			for len(stack) > 0 && stack[len(stack)-1] >= '0' && stack[len(stack)-1] <= '9' {
				count += int(stack[len(stack)-1]-'0') * i
				i *= 10
				stack = stack[:len(stack)-1]
			}

			res := ""
			for j := 0; j < count; j++ {
				res += stmp
			}
			stack = append(stack, res...)
		}
	}

	return string(stack)
}
