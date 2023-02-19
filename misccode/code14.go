package misccode

func CanPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%k != 0 {
		return false
	}

	r := make([]int, 0)
	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	process14(0, 0, sum/k, nums, r, &m)

	for _, v := range m {
		if v > 0 {
			return false
		}
	}
	return true
}
func process14(i, count, sum int, nums []int, r []int, m *map[int]int) {
	if count > sum {
		return
	}
	if i == len(nums) {
		if count == sum {
			for _, v := range r {
				(*m)[v]--
			}
		}
		return
	}
	r = append(r, nums[i])
	process14(i+1, count+nums[i], sum, nums, r, m)
	r = r[:len(r)-1]
	process14(i+1, count, sum, nums, r, m)
}
