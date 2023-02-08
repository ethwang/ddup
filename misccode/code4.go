package misccode

func ReversePairs(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	var pair int
	mergePair(nums, 0, len(nums)-1, &pair)
	return pair
}

func mergePair(nums []int, l, r int, pair *int) {
	if l == r {
		return
	}
	m := (l + r) / 2
	mergePair(nums, l, m, pair)
	mergePair(nums, m+1, r, pair)
	mergeReversePairs(nums, l, m, r, pair)
}

func mergeReversePairs(nums []int, l, m, r int, pair *int) {
	helps := make([]int, r-l+1)

	x := 0
	i := l
	j := m + 1
	for ; i <= m && j <= r; x++ {
		if nums[i] < nums[j] {
			helps[x] = nums[i]
			i++
		} else {
			*pair += j - m
			helps[x] = nums[j]
			j++
		}
	}

	for ; i <= m; x, i = x+1, i+1 {
		helps[x] = nums[i]
	}
	for ; j <= r; x, j = x+1, j+1 {
		helps[x] = nums[j]
	}

	for k, v := range helps {
		nums[l+k] = v
	}
}
