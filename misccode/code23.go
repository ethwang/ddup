package misccode

func topKFrequent(nums []int, k int) []int {

	m := make(map[int]int)
	for _, v := range nums {
		m[v]++
	}
	counts := make([][]int, 0)
	for key, val := range m {
		if len(counts) == k {
			if counts[0][1] < val {
				counts[0] = []int{key, val}
				heapify(counts, 0, k)
			}
			continue
		}
		counts = append(counts, []int{key, val})
		heapInsert(counts, len(counts)-1)
	}
	res := []int{}
	for i := 0; i < len(counts); i++ {
		res = append(res, counts[i][0])
	}
	return res
}
func heapInsert(counts [][]int, index int) {
	for counts[index][1] < counts[(index-1)/2][1] {
		counts[index], counts[(index-1)/2] = counts[(index-1)/2], counts[index]
		index = (index - 1) / 2
	}
}
func heapify(counts [][]int, index, size int) {
	left := index*2 + 1
	for left < size {
		least := left
		right := index*2 + 2
		if right < size {
			if counts[left][1] > counts[right][1] {
				least = right
			}
		}
		if counts[index][1] < counts[least][1] {
			least = index
		}
		if least == index {
			break
		}
		counts[index], counts[least] = counts[least], counts[index]
		index = least
		left = index*2 + 1
	}
}
