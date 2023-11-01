package misccode

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1, l2 := len(nums1), len(nums2)
	l := l1 + l2
	target1, target2 := -1, -1
	if l%2 == 0 {
		target1, target2 = l/2-1, l/2
	} else {
		target1 = l / 2
	}
	i, j, target := 0, 0, 0
	ans := make([]int, 0)
	for i < l1 && j < l2 {
		tmp := 0
		if nums1[i] < nums2[j] {
			tmp = nums1[i]
			i++
		} else {
			tmp = nums2[j]
			j++
		}
		if target == target1 {
			ans = append(ans, tmp)
		}
		if target == target2 {
			ans = append(ans, tmp)
		}
		target++
	}
	for ; i < l1; i++ {
		if target == target1 {
			ans = append(ans, nums1[i])
		}
		if target == target2 {
			ans = append(ans, nums1[i])
		}
		target++
	}
	for ; j < l2; j++ {
		if target == target1 {
			ans = append(ans, nums2[j])
		}
		if target == target2 {
			ans = append(ans, nums2[j])
		}
		target++
	}
	if len(ans) == 2 {
		return float64(ans[0]+ans[1]) / 2
	}
	if len(ans) == 1 {
		return float64(ans[0])
	}
	return 0
}
