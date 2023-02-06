package main

import (
	"c1/enumcode"
	"c1/greedycode"
	"c1/linklistcode"
	"c1/setcode"
	"c1/sortcode"
	"c1/treecode"
	"c1/trietreecode"
	"c1/utilcode"
	"fmt"
	"regexp"
	"sort"
)

// const regStr = "[:=, \r\n\t/]"

func RegFilter(str, regStr string) string {
	reg, _ := regexp.Compile(regStr)
	newStr := reg.ReplaceAllString(str, "")
	return newStr
}
func main() {
	// n皇后
	enumcode.Nqueen(4)
	return

	// 三数之和
	// fmt.Println(misccode.ThreeNumsSum([]int{-1, 0, 1, 2, -1, -4}, 0))

	// 字符串相同字符统计,空间复杂度O(1)
	s := "abbccc"
	st := enumcode.SameCharsCount(s)
	fmt.Println(st)
	return

	// 字符串子串
	s = "abcd"
	sss := enumcode.AllSubStrs(s)
	fmt.Println(sss)
	return

	// 字符串全排列
	s = "abc"
	enumcode.Permutations(s)
	return

	// 字符串子序列
	s = "abc"
	enumcode.AllSubsquences(s)
	return

	nums := []int{30, 1, 20, 4, 6}
	lessMoney := greedycode.LessMoneySplitGold(nums)
	fmt.Println(lessMoney)
	return

	testM := map[string]interface{}{
		"start_price":         1,
		"distance_unit_price": 4,
		"time_unit_price":     6,
	}
	type testT struct {
		StartPrice        int `json:"start_price"`
		DistanceUnitPrice int `json:"distance_unit_price"`
		TimeUnitPrice     int `json:"time_unit_price"`

		BeginTime string `json:"begin_time"`
		EndTime   string `json:"end_time"`
	}
	t := &testT{BeginTime: "111", EndTime: "222"}

	utilcode.MapToStruct(testM, t)

	fmt.Println(t)

	return

	set := setcode.NewSet()
	set.Add("test")
	fmt.Println(set.Has("test"))

	ss := []string{"jibw", "ji", "jp", "bw", "jibw"}
	res := greedycode.MinDicOrder(ss)
	fmt.Println(res)
	return

	ss = []string{"ba", "b"}
	res = greedycode.MinDicOrder(ss)
	fmt.Println(res)
	return

	ss = []string{"a", "ac", "ab"}
	res = greedycode.MinDicOrder(ss)
	fmt.Println(res)

	// trietreecode.TrieTree.Nodes = append(trietreecode.TrieTree.Nodes, &trietreecode.TrieTreeNode{})
	trietreecode.Insert("activity")
	trietreecode.Insert("act")
	trietreecode.Insert("acd")

	fmt.Printf("TrieTree: %v \n", trietreecode.PreCount("ac"))
	trietreecode.Del("act")
	fmt.Printf("TrieTree: %v \n", trietreecode.PreCount("ac"))
	trietreecode.Del("activity")
	fmt.Printf("TrieTree: %v \n", trietreecode.PreCount("ac"))

	return

	root := treecode.CreateCST(3)
	treecode.PrintCst(root)
	return

	head := linklistcode.InitRandomList([]int{3, 1, 7, 8, 6, 2, 4})
	headCopy := linklistcode.RandomNodeCopy(head)
	for head != nil && headCopy != nil {

		if head.Val != headCopy.Val || (head.Random != nil && headCopy.Random != nil && head.Random.Val != headCopy.Random.Val) {
			fmt.Println(false)
			break
		}
		head = head.Next
		headCopy = headCopy.Next
	}
	// return

	/*正则过滤*/
	//s := ":请=求,错/误\n\r\t fsfafdf"
	////	reg, _ := regexp.Compile(regStr)
	//ss := RegFilter(s, regStr)
	//fmt.Println(ss)
	//return

	/*链表回文判断*/
	//arr := []int{1}
	//h := linklistcode.InitLinkList(arr)
	//fmt.Println(linklistcode.IsPalindrome(h))
	//return

	/*排序*/
	arrTmp := utilcode.GenerateRandomArray(10, 30)
	fmt.Println(arrTmp)
	sortcode.Divide2(arrTmp, 5)
	fmt.Println(arrTmp)

	testTime := 500000
	maxSize := 10
	maxValue := 100
	succeed := true

	var arr1, arr2, arrOri []int
	var f func([]int)
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			fmt.Println("arrOri: ", arrOri)
		}
	}()
	for i := 0; i < testTime; i++ {
		arr1 = utilcode.GenerateRandomArray(maxSize, maxValue)
		arr2 = utilcode.CopyArray(arr1)
		arrOri = utilcode.CopyArray(arr1)
		//sortcode.InsertSort(arr1)
		// sortcode.SelectSort(arr1)
		// sortcode.MergeSort(arr1)
		//sortcode.HeapSort(arr1)
		// sortcode.BubbleSort(arr1)
		//f = sortcode.QuickSort
		//f(arr1)
		//f = sortcode.CountSort
		f = sortcode.RadixSort
		f(arr1)
		sort.Ints(arr2)
		if !utilcode.IsEqual(arr1, arr2) {
			succeed = false
			break
		}

		head := linklistcode.InitLinkList(arrOri)
		h := head
		for i := len(arrOri) - 1; i >= 0 && h != nil; i-- {

			if h.Value != arrOri[i] {
				succeed = false
				break
			}
			h = h.Next
		}

		// 反转
		rh := linklistcode.ReserveLinkList(head)
		for i := 0; i < len(arrOri) && rh != nil; i++ {

			if rh.Value != arrOri[i] {
				succeed = false
				break
			}
			rh = rh.Next
		}

		fmt.Println("------------")
		head = linklistcode.InitTwoWayLinkList(arrOri)
		fmt.Println(arrOri)
		h = head.Next
		var end *linklistcode.Node
		for h != nil {
			fmt.Printf("%v ", h.Value)
			end = h
			h = h.Next
		}
		fmt.Println()
		for end != nil && end != head {
			fmt.Printf("%v ", end.Value)
			end = end.Pre
		}

		fmt.Println("\n------------")
		newHead := linklistcode.ReserveTwoWayLinkList(head)
		nh := newHead
		var nhEnd *linklistcode.Node
		for nh != nil {
			fmt.Printf("%v ", nh.Value)
			nhEnd = nh
			nh = nh.Next
		}
		fmt.Println()
		for nhEnd != nil {
			fmt.Printf("%v ", nhEnd.Value)
			nhEnd = nhEnd.Pre
		}
	}
	if succeed {
		fmt.Println("Nice")
	} else {
		fmt.Println("Fucking")

		fmt.Println(arrOri)
		fmt.Println(arr1)
		fmt.Println(arr2)

		// 单步调试
		f(arrOri)
	}

}
