package utilcode

import (
	"math/rand"
	"time"
)

func GenerateRandomArray(maxSize, maxValue int) (arr []int) {
	rand.Seed(time.Now().Unix())
	arr = make([]int, rand.Intn(maxSize+1))
	for k := range arr {
		arr[k] = rand.Intn(maxValue + 1)
	}
	return arr
}

func CopyArray(arr []int) (arrCopy []int) {
	if arr == nil {
		return
	}
	arrCopy = make([]int, len(arr))
	copy(arrCopy, arr)
	return
}

func IsEqual(arr1, arr2 []int) bool {
	if (arr1 == nil && arr2 != nil) || (arr1 != nil && arr2 == nil) {
		return false
	}
	if arr1 == nil && arr2 == nil {
		return true
	}
	if len(arr1) != len(arr2) {
		return false
	}
	for k, v := range arr1 {
		if arr2[k] != v {
			return false
		}
	}
	return true

}
