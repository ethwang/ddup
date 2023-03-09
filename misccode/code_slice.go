package misccode

import "fmt"

func SliceT() {
	arr := make([]int, 3, 3) //创建一个长度为3，容量为4的切片
	fmt.Printf("%p\n", arr)  //0xc000012200
	// -----
	addNum(arr)
	// -----
	fmt.Printf("%p\n", arr) //0xc000012200
}

func addNum(sli []int) {
	sli = append(sli, 4)
	fmt.Printf("%p\n", sli) //0xc000012200
}
