package misccode

import "container/heap"

func MaxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	mh := &MaxH{}
	heap.Init(mh)
	for i := 0; i < k; i++ {
		heap.Push(mh, [2]int{nums[i], i})
	}
	ans = append(ans, (*mh)[0][0])

	for i := k; i < len(nums); i++ {
		heap.Push(mh, [2]int{nums[i], i})
		for len(*mh) > 0 && (*mh)[0][1] <= i-k {
			heap.Pop(mh)
		}
		ans = append(ans, (*mh)[0][0])
	}
	return ans
}

type MaxH [][2]int

func (mh MaxH) Less(i, j int) bool {
	return mh[i][0] > mh[j][0]
}
func (mh MaxH) Len() int {
	return len(mh)
}
func (mh MaxH) Swap(i, j int) {
	mh[i], mh[j] = mh[j], mh[i]
}
func (mh *MaxH) Push(x interface{}) {
	*mh = append(*mh, x.([2]int))
}
func (mh *MaxH) Pop() interface{} {
	x := (*mh)[len(*mh)-1]
	*mh = (*mh)[:len(*mh)-1]
	return x
}
