package misccode

func GenerateParenthesis(n int) []string {
	res := make([]string, 0)
	r := ""
	generate(n, n, r, &res)
	return res
	//curLen := len(res)
	//for i := 0; i < curLen; i++ {
	//	if IsValid(res[i]) {
	//		res = append(res, res[i])
	//	}
	//}
	//return res[curLen:]
}
func generate(lt, rt int, r string, res *[]string) {
	if lt < 0 || lt > rt {
		return
	}
	if lt == 0 && rt == 0 {
		*res = append(*res, r)
		return
	}
	r += "("
	generate(lt-1, rt, r, res)
	r = r[:len(r)-1]
	r += ")"
	generate(lt, rt-1, r, res)
}
