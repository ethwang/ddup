package misccode

func MaxProduct(nums []int) int {

	res := 0
	for i := 0; i < len(nums); i++ {
		tmpNum := nums[i]
		for j := i + 1; j < len(nums); j++ {
			tmpNum = nums[j] * tmpNum
			if tmpNum > res {
				res = tmpNum
			}
		}
	}
	return res
}
