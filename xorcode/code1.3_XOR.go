package xorcode

/*
	1.数列中有一个数字出现奇数次，返回该数字
	2.数列中有两个数字出现奇数次，返回这个两个数字
*/

func GetEvenTimesNum(nums []int) int {

	val := 0
	for i := 0; i < len(nums); i++ {
		val = val ^ nums[i]
	}
	return val
}

func GetEvenTimesNums(nums []int) (int, int) {

	val := 0
	valFlag := 0
	valB := 0

	for i := 0; i < len(nums); i++ {
		val = val ^ nums[i]
	}

	valFlag = val ^ (^val + 1)

	for i := 0; i < len(nums); i++ {

		if nums[i]&valFlag == 0 {
			valB = valB ^ nums[i]
		}
	}
	return valB, val ^ valB
}
