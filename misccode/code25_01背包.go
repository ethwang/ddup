package misccode

func MaxValue(weight, value []int, maxweight int) int {
	result := 0
	backtracking25(0, weight, value, 0, 0, maxweight, &result)
	return result
}

func backtracking25(i int, weight, value []int, curweight, curvalue, maxweigt int, result *int) {
	if curweight > maxweigt {
		return
	}
	if i == len(weight) {
		if curvalue > *result {
			*result = curvalue
		}
		return
	}
	backtracking25(i+1, weight, value, curweight+weight[i], curvalue+value[i], maxweigt, result)
	backtracking25(i+1, weight, value, curweight, curvalue, maxweigt, result)
}
