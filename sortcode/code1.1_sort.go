package sortcode

import (
	"math"
	"math/rand"
)

/*
	1.选择排序；
	2.冒泡排序；
	3.插入排序；
*/

// SelectSort 选择排序
func SelectSort(nums []int) {

	// 异常处理：数组长度=0 或 =1 直接返回
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	// 正常逻辑,每次找最小（最大）放到待排序的一个位置
	for i := 0; i < len(nums)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIndex] {
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

// BubbleSort 冒泡排序
func BubbleSort(nums []int) {

	// 异常处理：数组长度=0 或 =1 直接返回
	if len(nums) == 0 || len(nums) == 1 {
		return
	}

	// 正常逻辑
	for i := len(nums) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
}

// InsertSort 插入排序
func InsertSort(nums []int) {

	// 异常处理：数组长度=0 或 =1 直接返回
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	// 正常逻辑
	// 1,4   <-3
	// i=2,j=0 nums[i]=3,nums[j]=1
	// i=2,j=1 nums[i]=3,nums[j]=4 -->  1,3,4

	// 1,3,4 <-2
	// i=3,j=0 nums[i]=2,nums[j]=1
	// i=3,j=1 nums[i]=2,nums[j]=3 --> nums[i]=3,nums[j]=2 --> 1,2,4,3
	// i=3,j=2 nums[i]=3,nums[j]=4 --> nums[i]=4,nums[j]=3 --> 1,2,3,4
	for i := 1; i < len(nums); i++ {
		/// 错误代码
		//for j := 0; j < i && nums[j] > nums[i]; j++ {
		//	nums[i], nums[j] = nums[j], nums[i]
		//}
		for j := i - 1; j >= 0 && nums[j+1] < nums[j]; j-- {
			nums[j+1], nums[j] = nums[j], nums[j+1]
		}
	}
}

/*
	4.归并排序
	5.堆排序
	6.快速排序
*/

// MergeSort 归并排序
func MergeSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	mergeSort(nums, 0, len(nums)-1)
}

func mergeSort(nums []int, l, r int) {
	// 递归终止条件
	if l == r {
		return
	}
	m := (l + r) / 2
	mergeSort(nums, l, m)
	mergeSort(nums, m+1, r)
	merge(nums, l, m, r)
}

func merge(nums []int, l, m, r int) {
	help := make([]int, r-l+1)
	i := 0
	l0, m0 := l, m+1

	for ; l0 <= m && m0 <= r; i++ {

		if nums[l0] < nums[m0] {
			help[i] = nums[l0]
			l0++
		} else {
			help[i] = nums[m0]
			m0++
		}
	}
	for ; l0 <= m; i, l0 = i+1, l0+1 {
		help[i] = nums[l0]
	}
	for ; m0 <= r; i, m0 = i+1, m0+1 {
		help[i] = nums[m0]
	}
	for k, v := range help {
		nums[l+k] = v
	}
}

// 归并扩展：小和问题 和 逆序对
// MergeSmallSum 如果不用map,怎么求解?
func MergeSmallSumAndReverseOrderPair(nums []int) (smallSum map[int]int, reverseOrderPair map[int][]int) {
	smallSum, reverseOrderPair = make(map[int]int), make(map[int][]int)
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	mergeSmallSumAndReverseOrderPair(nums, 0, len(nums)-1, smallSum, reverseOrderPair)
	return
}

func mergeSmallSumAndReverseOrderPair(nums []int, l, r int, smallSum map[int]int, reverseOrderPair map[int][]int) {
	if l == r {
		return
	}
	m := (l + r) / 2
	mergeSmallSumAndReverseOrderPair(nums, l, m, smallSum, reverseOrderPair)
	mergeSmallSumAndReverseOrderPair(nums, m+1, r, smallSum, reverseOrderPair)
	smallSumAndReverseOrderPair(nums, l, m, r, smallSum, reverseOrderPair)
}

func smallSumAndReverseOrderPair(nums []int, l, m, r int, smallSum map[int]int, reverseOrderPair map[int][]int) {
	help := make([]int, r-l+1)
	i := 0
	l0, m0 := l, m+1

	for ; l0 <= m && m0 <= r; i++ {

		if nums[l0] < nums[m0] {
			// 小和
			smallSum[nums[l0]] += r - m0 + 1
			help[i] = nums[l0]
			l0++
		} else {
			//  逆序对
			for k := l0; k <= m; k++ {
				reverseOrderPair[nums[k]] = append(reverseOrderPair[nums[k]], nums[m0])
			}
			help[i] = nums[m0]
			m0++
		}
	}
	for ; l0 <= m; i, l0 = i+1, l0+1 {
		help[i] = nums[l0]
	}
	for ; m0 <= r; i, m0 = i+1, m0+1 {
		help[i] = nums[m0]
	}
	for k, v := range help {
		nums[l+k] = v
	}
}

// HeapSort 堆排序
func HeapSort(nums []int) {

	for k := range nums {
		heapInsert(nums, k)
	}
	size := len(nums) - 1
	nums[size], nums[0] = nums[0], nums[size]

	for size > 0 {
		heapify(nums, 0, size)
		size--
		nums[size], nums[0] = nums[0], nums[size]
	}

	/*
		// 堆排序为啥不用这种方式？ 分析时间复杂度O(n^2)和空间复杂度
		for j := len(nums) - 1; j >= 0; j-- {

			for i := 0; i <= j; i++ {
				heapInsert(nums, i)
			}
			nums[j], nums[0] = nums[0], nums[j]
		}
	*/
}

// 大顶堆
func heapInsert(nums []int, index int) {

	for nums[index] > nums[(index-1)/2] {
		nums[index], nums[(index-1)/2] = nums[(index-1)/2], nums[index]
		index = (index - 1) / 2
	}
}

// 大顶堆
func heapify(nums []int, index, size int) {
	left := 2*index + 1

	for left < size {
		largest := left

		right := left + 1
		if right < size {
			if nums[right] > nums[left] {
				largest = right
			}
		}
		if nums[index] > nums[largest] {
			largest = index
		}
		if largest == index {
			break
		}
		nums[largest], nums[index] = nums[index], nums[largest]
		index = largest
		left = 2*index + 1
	}

}

// 引入快速排序，荷兰国旗问题
func Divide(nums []int, num int) {

	// 方法1
	for i := 0; i < len(nums); i++ {
		if nums[i] > num {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] < num {
					nums[j], nums[i] = nums[i], nums[j]
					break
				}
			}
		}
	}
	// 方法2
	// 核心思想：划分区域,定义一个大于等于区域的索引范围
	more := len(nums) - 1
	for i := 0; i <= more; {
		if nums[i] >= num {
			nums[i], nums[more] = nums[more], nums[i]
			more--
		} else {
			i++
		}
	}

}
func Divide2(nums []int, num int) {

	less := -1
	more := len(nums)
	index := 0

	// 核心思想：划分区域，划分为小于，等于，大于区域
	for index < more {
		if nums[index] < num {
			less++
			nums[less], nums[index] = nums[index], nums[less]
			index++
		} else if nums[index] > num {
			more--
			nums[index], nums[more] = nums[more], nums[index]
		} else {
			index++
		}
	}
}

// QuickSort快速排序
func QuickSort(nums []int) {
	if len(nums) == 0 || len(nums) == 1 {
		return
	}
	quickSort(nums, 0, len(nums)-1)
}

func quickSort(nums []int, l, r int) {
	if l >= r {
		return
	}
	// 随机选择一个数做为划分值
	randomPosition := l + rand.Intn(r-l+1)
	nums[randomPosition], nums[r] = nums[r], nums[randomPosition]
	pl, pr := partition(nums, l, r)
	quickSort(nums, l, pl)
	quickSort(nums, pr+1, r)
}

func partition(nums []int, l, r int) (int, int) {
	less, more := l-1, r

	for l < more {
		if nums[l] > nums[r] {

			more--
			nums[l], nums[more] = nums[more], nums[l]
		} else if nums[l] < nums[r] {
			less++
			nums[less], nums[l] = nums[l], nums[less]
			l++
		} else {
			l++
		}
	}
	/*
		/// 错误答案
			for i := 0; i < more; {
				if nums[i] > nums[r] {
					more--
					nums[more], nums[i] = nums[i], nums[more]
				} else if nums[i] < nums[r] {
					less++
					nums[i], nums[less] = nums[less], nums[i]
					i++
				} else {
					i++
				}
			}
	*/
	nums[more], nums[r] = nums[r], nums[more]

	return less, more
}

/*
	7.计数排序
	8.基数排序
*/
// CountSort计数排序
func CountSort(nums []int) {

	if nums == nil || len(nums) <= 1 {
		return
	}
	scope := math.MinInt32
	for _, v := range nums {
		scope = int(math.Max(float64(v), float64(scope)))
	}
	newNums := make([]int, scope+1)
	for _, v := range nums {
		newNums[v]++
	}

	i := 0
	for k, v := range newNums {
		for ; v > 0; v-- {
			nums[i] = k
			i++
		}
	}
}

// RadixSort 基数排序
func RadixSort(nums []int) {

	if nums == nil || len(nums) <= 1 {
		return
	}
	radixSort(nums, 0, len(nums)-1)
}
func radixSort(nums []int, begin, end int) {
	digits := getMaxDigits(nums)

	tmp := make([]int, end-begin+1)
	for i := 1; i <= digits; i++ {

		vds := make([]int, 10)
		for j := begin; j <= end; j++ {
			vd := getValueWithDigit(nums[j], i)
			vds[vd]++
		}
		for k := 1; k < len(vds); k++ {
			vds[k] += vds[k-1]
		}
		//for l := begin; l <= end; l++ {

		//	vd := getValueWithDigit(nums[l], i)
		//	tmp[vds[vd]-1] = nums[l]
		//	vds[vd]--

		//}
		for l := end; l >= begin; l-- {
			vd := getValueWithDigit(nums[l], i)
			tmp[vds[vd]-1] = nums[l]
			vds[vd]--
		}
		for m, n := begin, 0; m <= end; m, n = m+1, n+1 {
			nums[m] = tmp[n]
		}
	}

}
func getValueWithDigit(num, digit int) int {

	nd := 0
	for digit > 0 {
		nd = num % 10
		num = num / 10
		digit--
	}
	return nd
}

func getMaxDigits(nums []int) int {

	digits := 0
	maxDigitsValue := nums[0]
	for i := 1; i < len(nums); i++ {

		if nums[i] > maxDigitsValue {
			maxDigitsValue = nums[i]
		}
	}
	for maxDigitsValue > 0 {
		digits++
		maxDigitsValue = maxDigitsValue / 10
	}
	return digits
}
