package greedycode

import "sort"

func LessMoneySplitGold(golds []int) (lessMoney int) {

	goldLen := 0
	for _, v := range golds {
		goldLen += v
	}
	lessMoney = goldLen
	sort.Ints(golds)

	lessMoney = goldLen
	curLess := goldLen

	for i := len(golds) - 1; i > 1; i-- {
		lessMoney += curLess - golds[i]
		curLess = curLess - golds[i]
	}

	return lessMoney
}
