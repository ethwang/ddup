package misccode

import (
	"container/heap"
	"sort"
)

/*
维护两个元素数量相差小于等于1堆: 小顶堆和大顶堆;

	如果数量相同: median=(两堆顶之和)/2; 如果堆顶元素数量相差1: median=元素数量多的堆的堆顶元素;
*/
type MedianFinder struct {
	H1    []int
	H2    []int
	H1Max int
	H2Min int
}

func ConstructorMedianFinder() MedianFinder {
	return MedianFinder{
		H1:    []int{},
		H2:    []int{},
		H1Max: 0,
		H2Min: 0,
	}
}

func (this *MedianFinder) AddNum(num int) {
	if num < this.H1Max {
		if len(this.H1) <= len(this.H2) {
			this.H1 = append(this.H1, num)
			this.H1Max = findArrayMax(this.H1)
		} else {
			// 将当前H1的最大值放到H2中重新找H2的最小值, 将num加入H1中重新找H1的最大值
			this.H2 = append(this.H2, this.H1Max)
			this.H2Min = findArrayMin(this.H2)

			this.H1 = delArrayVal(this.H1, this.H1Max)
			this.H1 = append(this.H1, num)
			this.H1Max = findArrayMax(this.H1)
		}

	} else if num >= this.H1Max && num <= this.H2Min {
		if len(this.H1) <= len(this.H2) {
			this.H1 = append(this.H1, num)
			this.H1Max = findArrayMax(this.H1)
		} else {
			this.H2 = append(this.H2, num)
			this.H2Min = findArrayMin(this.H2)
		}
	} else {
		if len(this.H2) <= len(this.H1) {
			this.H2 = append(this.H2, num)
			this.H2Min = findArrayMin(this.H2)
		} else {
			// 将当前H2的最小值放到H1中重新找H1的最大值, 将num加入H2中重新找H2的最小值
			this.H1 = append(this.H1, this.H2Min)
			this.H1Max = findArrayMax(this.H1)

			this.H2 = delArrayVal(this.H2, this.H2Min)
			this.H2 = append(this.H2, num)
			this.H2Min = findArrayMin(this.H2)
		}
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if len(this.H1) == len(this.H2) {
		return float64(this.H1Max+this.H2Min) / 2
	} else if len(this.H1) > len(this.H2) {
		return float64(this.H1Max)
	} else {
		return float64(this.H2Min)
	}
}

func delArrayVal(nums []int, val int) []int {
	for i := 0; i < len(nums); i++ {
		if nums[i] == val {
			if len(nums) == 1 {
				return []int{}
			}
			if i < len(nums)-1 {
				return append(nums[:i], nums[i+1:]...)
			} else {
				return nums[:i]
			}
		}
	}
	return nums
}
func findArrayMin(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	min := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
	}
	return min
}
func findArrayMax(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	max := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	return max
}

//type MinHeap struct {
//	H    []int
//	Size int
//}
//
//func (mh *MinHeap) Pop() {
//
//}
//
///* 小顶堆 */
//func (mh *MinHeap) MinHeapInsert(target int) {
//	// 使用size控制
//	mh.H = append(mh.H, target)
//	mh.Size++
//	size := mh.Size
//	bottom := size - 1
//	for mh.H[bottom] < mh.H[(bottom-1)/2] {
//		mh.H[bottom], mh.H[(bottom-1)/2] = mh.H[(bottom-1)/2], mh.H[bottom]
//		bottom = (bottom - 1) / 2
//	}
//}
//func (mh *MinHeap) MinHeapify(top int) {
//	left := 2*top + 1
//	for mh.Size > left {
//		largestIndex := left
//		if mh.Size > left+1 {
//			if mh.H[left+1] < mh.H[left] {
//				largestIndex = left + 1
//			}
//		}
//		if mh.H[top] < mh.H[largestIndex] {
//			return
//		}
//		mh.H[top], mh.H[largestIndex] = mh.H[largestIndex], mh.H[top]
//		top = largestIndex
//		left = 2*top + 1
//	}
//}
//
//type MaxHeap struct {
//	H    []int
//	Size int
//}
//
///* 大顶堆 */
//func (mh *MaxHeap) MaxHeapInsert(target int) {
//	mh.H = append(mh.H, target)
//	mh.Size++
//	bottom := len(mh.H) - 1
//	//nums = append(nums, target)
//	//bottom := len(nums) - 1
//	for mh.H[bottom] > mh.H[(bottom-1)/2] {
//		mh.H[bottom], mh.H[(bottom-1)/2] = mh.H[(bottom-1)/2], mh.H[bottom]
//		bottom = (bottom - 1) / 2
//	}
//}
//func (mh *MaxHeap) MaxHeapify(top int) {
//	left := 2*top + 1
//	for mh.Size > left {
//		largestIndex := left
//		if mh.Size >= left+1 {
//			if mh.H[left+1] > mh.H[left] {
//				largestIndex = left + 1
//			}
//		}
//		if mh.H[top] > mh.H[largestIndex] {
//			return
//		}
//		mh.H[top], mh.H[largestIndex] = mh.H[largestIndex], mh.H[top]
//		top = largestIndex
//		left = 2*top + 1
//	}
//}
//
//type MedianFinder2 struct {
//	HMin *MinHeap
//	HMax *MaxHeap
//}
//
//func ConstructorMedianFinder2() MedianFinder2 {
//	return MedianFinder2{
//		HMin: &MinHeap{Size: 0, H: []int{}},
//		HMax: &MaxHeap{Size: 0, H: []int{}},
//	}
//}
//
//func (this *MedianFinder2) AddNum(num int) {
//	if len(this.HMax.H) == 0 {
//		this.HMax.MaxHeapInsert(num)
//	}
//	if len(this.HMin.H) == 0 {
//		this.HMin.MinHeapInsert(num)
//	}
//	if num < this.HMax.H[0] {
//		if this.HMax.Size <= this.HMin.Size {
//			this.HMax.MaxHeapInsert(num)
//		} else {
//			this.HMin.MinHeapInsert(this.HMax.H[0])
//
//			this.HMax.H[0], this.HMax.H[this.HMax.Size-1] = this.HMax.H[this.HMax.Size-1], this.HMax.H[0]
//			this.HMax.Size--
//
//		}
//
//	} else if num >= this.HMax.H[0] && num <= this.HMin.H[0] {
//
//	} else {
//
//	}
//}
//
//func (this *MedianFinder2) FindMedian() float64 {
//}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */

// An IntHeap is a min-heap of ints.
//type IntHeap []int
//
//func (h IntHeap) Len() int           { return len(h) }
//func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小顶堆
//func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
//
//func (h *IntHeap) Push(x interface{}) {
//	// Push and Pop use pointer receivers because they modify the slice's length,
//	// not just its contents.
//	*h = append(*h, x.(int))
//}
//
//func (h *IntHeap) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//
//// This example inserts several ints into an IntHeap, checks the minimum,
//// and removes them in order of priority.
//func main() {
//	h := &IntHeap{2, 1, 5}
//	heap.Init(h)
//	heap.Push(h, 3)
//	fmt.Printf("minimum: %d\n", (*h)[0]) // minimum: 1
//	for h.Len() > 0 {
//		fmt.Printf("%d ", heap.Pop(h)) // 1 2 3 5
//	}
//}

type HeapInterface interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
	Push(x interface{})
	Pop() interface{}
	Top() interface{}
}

func up(hi HeapInterface, i int) {
	for {
		j := (i - 1) / 2 // parent
		if i == j || !hi.Less(i, j) {
			break
		}
		hi.Swap(i, j)
		i = j
	}
}

func down(hi HeapInterface, i, n int) bool {
	i0 := i
	for {
		j := 2*i + 1
		if j >= n || j < 0 {
			break
		}
		j0 := j
		if j+1 < n && hi.Less(j+1, j) {
			j0 = j + 1
		}
		if hi.Less(i, j0) {
			break
		}
		hi.Swap(i, j0)
		i = j0
	}
	return i > i0
}

func Push(hi HeapInterface, val interface{}) {
	hi.Push(val)
	up(hi, hi.Len()-1)
}

func Pop(hi HeapInterface) interface{} {
	n := hi.Len() - 1
	hi.Swap(0, n)
	down(hi, 0, n)
	return hi.Pop()
}
func Init(hi HeapInterface) {
	n := hi.Len()
	for i := n/2 - 1; i >= 0; i-- {
		down(hi, i, n)
	}
}

type MinHeap []int

func (mh MinHeap) Len() int {
	return len(mh)
}
func (mh MinHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh MinHeap) Less(i, j int) bool {
	return mh[i] < mh[j]
}
func (mh *MinHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}
func (mh *MinHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}
func (mh MinHeap) Top() interface{} {
	if mh.Len() > 0 {
		return mh[0]
	}
	return nil
}

type MaxHeap []int

func (mh MaxHeap) Len() int {
	return len(mh)
}
func (mh MaxHeap) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh MaxHeap) Less(i, j int) bool {
	return mh[i] > mh[j]
}
func (mh *MaxHeap) Push(x interface{}) {
	*mh = append(*mh, x.(int))
}
func (mh *MaxHeap) Pop() interface{} {
	old := *mh
	n := len(old)
	x := old[n-1]
	*mh = old[0 : n-1]
	return x
}
func (mh MaxHeap) Top() interface{} {
	if mh.Len() > 0 {
		return mh[0]
	}
	return nil
}

type MedianFinder2 struct {
	MaxH MaxHeap
	MinH MinHeap
}

func ConstructorMedianFinder2() MedianFinder2 {
	maxH := MaxHeap{}
	minH := MinHeap{}
	Init(&maxH)
	Init(&minH)
	return MedianFinder2{
		MaxH: maxH,
		MinH: minH,
	}
}
func (this *MedianFinder2) AddNum(num int) {
	if this.MaxH.Len() == 0 || num <= this.MaxH.Top().(int) {
		Push(&this.MaxH, num)
		if this.MaxH.Len() > this.MinH.Len()+1 {
			Push(&this.MinH, Pop(&this.MaxH).(int))
		}
	} else {
		Push(&this.MinH, num)
		if this.MinH.Len() > this.MaxH.Len() {
			Push(&this.MaxH, Pop(&this.MinH).(int))
		}
	}

	//	if this.MaxH.Top() == nil {
	//		Push(&this.MaxH, num)
	//		return
	//	}
	//
	//	if this.MinH.Top() == nil {
	//		Push(&this.MinH, num)
	//		return
	//	}
	//
	//	if this.MaxH.Len() == 1 && this.MinH.Len() == 1 {
	//		if this.MaxH.Top().(int) > this.MinH.Top().(int) {
	//			tmp := this.MaxH.Pop()
	//			this.MaxH.Push(this.MinH.Pop())
	//			this.MinH.Push(tmp)
	//		}
	//		return
	//	}
	//
	//	if num < this.MaxH.Top().(int) {
	//		if this.MaxH.Len() <= this.MinH.Len() {
	//			Push(&this.MaxH, num)
	//		} else {
	//			Push(&this.MinH, Pop(&this.MaxH))
	//			Push(&this.MaxH, num)
	//		}
	//	} else if num >= this.MaxH.Top().(int) && num <= this.MinH.Top().(int) {
	//
	//		if this.MaxH.Len() <= this.MinH.Len() {
	//			Push(&this.MaxH, num)
	//		} else {
	//			Push(&this.MinH, num)
	//		}
	//	} else {
	//
	//		if this.MinH.Len() <= this.MaxH.Len() {
	//			Push(&this.MinH, num)
	//		} else {
	//			Push(&this.MaxH, Pop(&this.MinH))
	//			Push(&this.MinH, num)
	//		}
	//	}
}
func (this *MedianFinder2) FindMedian() float64 {
	if this.MaxH.Len() > this.MinH.Len() {
		return float64(this.MaxH.Top().(int))
	}
	return (float64(this.MaxH.Top().(int)) + float64(this.MinH.Top().(int))) / 2
	//if this.MaxH.Len() == this.MinH.Len() {
	//	if this.MaxH.Top() != nil && this.MinH.Top() != nil {
	//		return (float64(this.MaxH.Top().(int)) + float64(this.MinH.Top().(int))) / 2
	//	}
	//} else if this.MaxH.Len() > this.MinH.Len() {
	//	return float64(this.MaxH.Top().(int))
	//} else {
	//	return float64(this.MinH.Top().(int))
	//}
	//return -1
}

type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

type MedianFinder3 struct {
	MaxH, MinH *hp
}

func (this *MedianFinder3) AddNum(num int) {
	maxH, minH := this.MaxH, this.MinH

	if maxH.Len() == 0 || num <= -maxH.IntSlice[0] {
		heap.Push(maxH, -num)
		if maxH.Len() > minH.Len()+1 {
			heap.Push(minH, -heap.Pop(maxH).(int))
		}
	} else {
		heap.Push(minH, num)
		if minH.Len() > maxH.Len() {
			heap.Push(maxH, -heap.Pop(minH).(int))
		}
	}
}

func (this *MedianFinder3) FindMedian() float64 {
	maxH, minH := this.MaxH, this.MinH
	if maxH.Len() > minH.Len() {
		return float64(maxH.IntSlice[0])
	}
	return (float64(maxH.IntSlice[0]) + float64(minH.IntSlice[0])) / 2
}
