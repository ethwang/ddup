package enumcode

import (
	"fmt"
	"math"
)

func Nqueen(n int) {

	arr := make([]int, n)
	res := make([][]string, 0)
	process7(0, arr, &res)
	fmt.Println(res)
	// fmt.Println(result)
}
func genBoard(arr []int, n int) []string {

	res := make([]string, 0)
	for i := 0; i < n; i++ {
		column := make([]byte, n)
		for j := 0; j < n; j++ {
			column[j] = '.'
		}
		column[arr[i]] = 'Q'
		res = append(res, string(column))
	}
	return res
}

func process7(i int, arr []int, res *[][]string) {

	if i == len(arr) {
		dst := make([]int, len(arr))
		copy(dst, arr)
		*res = append(*res, genBoard(dst, len(arr)))
		return
	}

	for j := 0; j < len(arr); j++ {
		if isValid(arr, i, j) {
			arr[i] = j
			process7(i+1, arr, res)
			// 理论上是应该重置为初始值的，但是因为isValid方法只比较当前行之前的点，所以可以不置为初始值
			// arr[i] = -1
		}
	}
}
func isValid(arr []int, i, j int) bool {

	// 比较i行之前的行
	for k := 0; k < i; k++ {

		if arr[k] == j || math.Abs(float64(k-i)) == math.Abs(float64(arr[k]-j)) {
			return false
		}
	}
	return true
}

func process8(i int, arr []int) int {

	res := 0
	if i == len(arr) {
		return 1
	}

	for j := 0; j < len(arr); j++ {
		if isValid(arr, i, j) {
			arr[i] = j
			res += process8(i+1, arr)
			// 理论上是应该重置为初始值的，但是因为isValid方法只比较当前行之前的点，所以可以不置为初始值
			// arr[i] = -1
		}
	}
	return res
}
