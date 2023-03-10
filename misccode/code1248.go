package misccode

func NumberOfSubarrays(nums []int, k int) int {
	// map[k]v nums前i个数字中有k个奇数，v=有k个奇数的子字符串个数
	m := map[int]int{0: 1}

	// 奇数个数
	ncount := 0
	// 优美子数组个数
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 1 {
			ncount++
		}
		if m[ncount-k] > 0 {
			count += m[ncount-k]
		}
		m[ncount]++
	}
	return count
}
