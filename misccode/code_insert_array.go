package misccode

func InsertArray(slice []int, index int, value int) []int {
	position := index - 1
	if position > len(slice) {
		position = len(slice)
	}
	// 插入元素，golang没有直接插入到切片的方法，可以通过下面的方式实现
	slice = append(slice[:position], append([]int{value}, slice[position:]...)...)
	return slice
}

//func main() {
//    s := []int{1, 2, 3, 4, 5}
//    n := 3 // 要插入的位置
//    value := 6 // 要插入的值
//    s = insert(s, n, value)
//    fmt.Println(s) // 输出: [1 2 3 6 4 5]
//}
