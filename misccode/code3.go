package misccode

func mergeSort(nums []int, l, r int) {
	if l == r {
		return
	}

	m := (l + r) / 2
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, r)
	merge(nums, l, m, r)
}
func merge(nums []int, l, m, r int) {

	tmp := make([]int, r-l+1)
	x := 0
	i := l
	j := m + 1
	for ; i <= m && j <= r; x++ {

		if nums[i] < nums[j] {
			tmp[x] = nums[i]
			i++
		} else {
			tmp[x] = nums[j]
			j++
		}
	}
	for ; i <= m; i, x = i+1, x+1 {
		tmp[x] = nums[i]
	}
	for ; j <= r; j, x = j+1, x+1 {
		tmp[x] = nums[j]
	}
	for m := 0; m < len(tmp); m++ {
		nums[m+l] = tmp[m]
	}
}
