package misccode

import "sort"

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

func CanPartitionKSubsets14(nums []int, k int) bool {

	total := 0
	for _, v := range nums {
		total += v
	}
	if total%k != 0 {
		return false
	}
	sum := total / k
	path := make([]int, 0)
	result := make([]int, 0)

	backtracking14(0, nums, 0, sum, path, &result)

	sort.Ints(nums)
	sort.Ints(result)
	if len(nums) != len(result) {
		return false
	}
	for i := 0; i < len(nums); i++ {
		if nums[i] != result[i] {
			return false
		}
	}
	return true
}

func backtracking14(startIndex int, nums []int, count int, sum int, path []int, result *[]int) {
	if count == sum {
		dst := make([]int, len(path))
		copy(dst, path)
		*result = append(*result, dst...)
		return
	}
	m := make(map[int]bool)
	for i := startIndex; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		if count+nums[i] > sum {
			continue
		}
		m[nums[i]] = true
		count += nums[i]
		path = append(path, nums[i])
		backtracking14(i+1, nums, count, sum, path, result)
		count -= nums[i]
		path = path[:len(path)-1]
	}
}
func CanPartitionKSubsets141(nums []int, k int) bool {

	total := 0
	for _, v := range nums {
		total += v
	}
	if total%k != 0 {
		return false
	}
	sum := total / k
	bucket := make([]int, k)

	return backtracking141(0, nums, bucket, sum, k)
}
func backtracking141(index int, nums []int, bucket []int, sum int, k int) bool {
	if index == len(nums) {
		return true
	}
	for i := 0; i < k; i++ {
		if bucket[i]+nums[index] > sum {
			continue
		}
		bucket[i] += nums[index]
		if backtracking141(index+1, nums, bucket, sum, k) {
			return true
		}
		bucket[i] -= nums[index]
	}
	return false
}
