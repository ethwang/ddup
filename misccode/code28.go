package misccode

import "sort"

func ReconstructQueue(people [][]int) [][]int {
	queue := make([][]int, len(people))
	sort.Slice(people, func(i, j int) bool {

		if people[i][0] < people[j][0] {
			return true
		} else if people[i][0] == people[j][0] {
			if people[i][1] > people[j][1] {
				return true
			}
		}
		return false

	})

	for i := 0; i < len(people); i++ {
		p := people[i][1] + 1
		for j := 0; j < len(queue); j++ {
			if queue[j] == nil {
				p--
			}
			if p == 0 {
				queue[j] = people[i]
				break
			}
		}
	}
	return queue
}
