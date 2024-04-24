package main

import (
	"c1/enumcode"
	"c1/greedycode"
	"c1/linklistcode"
	"c1/misccode"
	"c1/queuecode"
	"c1/setcode"
	"c1/sortcode"
	"c1/treecode"
	"c1/trietreecode"
	"c1/utilcode"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"regexp"
	"sort"
	"strings"
<<<<<<< Updated upstream
=======
	"sync"
	"syscall"
>>>>>>> Stashed changes
	"time"
	"unsafe"

	"context"
)

// const regStr = "[:=, \r\n\t/]"

func RegFilter(str, regStr string) string {
	reg, _ := regexp.Compile(regStr)
	newStr := reg.ReplaceAllString(str, "")
	return newStr
}
func test() (i int) {
	defer func() {
		i++
	}()
	return
}
func test2() int {
	var i int
	defer func() {
		i++
	}()

	return i
}

type FamilyTaskBoxGift struct {
	GiftImg   string `json:"gift_img"`
	GiftName  string `json:"gift_name"`
	GiftAging string `json:"gift_aging"`
	GiftCount int32  `json:"gift_count"`
}

//	type FamilyBoxCoinsT struct {
//		FamilyBoxCoins map[int]map[int]int `json:"family_box_coins"`
//	}
type FamilyBoxCoinsT map[int]map[int]int
type FamilyBoxActiveT map[int]map[int]int
type FamilyBoxResourcesT map[int]map[int][]*Resource

var FamilyBoxCoinss = FamilyBoxCoinsT{
	1:  {1: 1, 2: 3, 3: 6, 4: 10},
	2:  {1: 3, 2: 6, 3: 10, 4: 15},
	3:  {1: 5, 2: 10, 3: 14, 4: 18},
	4:  {1: 7, 2: 12, 3: 18, 4: 24},
	5:  {1: 8, 2: 15, 3: 20, 4: 26},
	6:  {1: 10, 2: 20, 3: 24, 4: 28},
	7:  {1: 12, 2: 26, 3: 30, 4: 36},
	8:  {1: 14, 2: 28, 3: 32, 4: 38},
	9:  {1: 16, 2: 32, 3: 36, 4: 40},
	10: {1: 17, 2: 36, 3: 44, 4: 45},
	11: {1: 18, 2: 38, 3: 46, 4: 52},
	12: {1: 20, 2: 40, 3: 48, 4: 60},
	13: {1: 22, 2: 45, 3: 50, 4: 70},
	14: {1: 24, 2: 50, 3: 60, 4: 80},
	15: {1: 26, 2: 65, 3: 90, 4: 100},
}

var FamilyBoxActive = FamilyBoxActiveT{
	1:  {1: 525, 2: 2100, 3: 5250, 4: 9800},
	2:  {1: 1680, 2: 5670, 3: 10500, 4: 17640},
	3:  {1: 3360, 2: 10920, 3: 19600, 4: 31360},
	4:  {1: 5600, 2: 17850, 3: 31500, 4: 49000},
	5:  {1: 7560, 2: 23940, 3: 42000, 4: 64680},
	6:  {1: 9800, 2: 30870, 3: 53900, 4: 82320},
	7:  {1: 12320, 2: 38640, 3: 67200, 4: 101920},
	8:  {1: 15120, 2: 47250, 3: 81900, 4: 123480},
	9:  {1: 18200, 2: 56700, 3: 98000, 4: 147000},
	10: {1: 20020, 2: 62370, 3: 115500, 4: 172480},
	11: {1: 23520, 2: 73080, 3: 134400, 4: 211680},
	12: {1: 25480, 2: 81900, 3: 145600, 4: 229320},
	13: {1: 31500, 2: 100800, 3: 178500, 4: 294000},
	14: {1: 36750, 2: 132300, 3: 245000, 4: 377300},
	15: {1: 44800, 2: 168000, 3: 350000, 4: 548800},
}

type Resource struct {
	Type  int32  `json:"type"`  // 资源类型 1 礼物； 2 头像框； 3 勋章； 4 座驾； 5 称号； 6 聊天气泡； 7 迷你卡
	Id    int32  `json:"id"`    // 资源id
	Img   string `json:"img"`   // 资源图片
	Term  int64  `json:"term"`  // 资源期限,单位/天
	Count int64  `json:"count"` // 资源总数量

}

var FamilyBoxResources = FamilyBoxResourcesT{
	1: {1: []*Resource{},
		2: []*Resource{},
		3: []*Resource{},
		4: []*Resource{},
	},
	2: {1: []*Resource{},
		2: []*Resource{},
		3: []*Resource{},
		4: []*Resource{},
	},
	3: {1: []*Resource{},
		2: []*Resource{},
		3: []*Resource{},
		4: []*Resource{{Id: 64, Img: "media/2023/06/26/e8328bda140c11eeb2540242c0a80002.png", Type: 6, Term: 1, Count: 5}},
	},
	4: {1: []*Resource{},
		2: []*Resource{},
		3: []*Resource{{Id: 65, Img: "media/2023/06/26/fb29c65e140c11eeb3250242c0a80002.png", Type: 6, Term: 1, Count: 8}},
		4: []*Resource{{Id: 396, Img: "media/2023/02/06/40c09010a5f611ed8b8e0242c0a80002.png", Type: 2, Term: 1, Count: 8}},
	},
	5: {
		1: []*Resource{},
		2: []*Resource{},
		3: []*Resource{{Id: 65, Img: "media/2023/06/26/fb29c65e140c11eeb3250242c0a80002.png", Type: 6, Term: 1, Count: 10}},
		4: []*Resource{{Id: 396, Img: "media/2023/02/06/40c09010a5f611ed8b8e0242c0a80002.png", Type: 2, Term: 1, Count: 10}},
	},
	6: {
		1: []*Resource{},
		2: []*Resource{},
		3: []*Resource{{Id: 48, Img: "media/2023/04/20/3bdd1880df7a11edb2a00242c0a80002.png", Type: 6, Term: 1, Count: 10}},
		4: []*Resource{{Id: 9, Img: "media/2023/06/27/e64123a414b611ee8c170242c0a80002.png", Type: 2, Term: 1, Count: 10}},
	},
	7: {
		1: []*Resource{},
		2: []*Resource{{Id: 48, Img: "media/2023/04/20/3bdd1880df7a11edb2a00242c0a80002.png", Type: 6, Term: 1, Count: 10}},
		3: []*Resource{{Id: 9, Img: "media/2023/06/27/e64123a414b611ee8c170242c0a80002.png", Type: 2, Term: 1, Count: 10}},
		4: []*Resource{{Id: 191, Img: "media/2022/06/20/0dc5d982f08811ec89fd0242c0a80002.png", Type: 2, Term: 1, Count: 10}},
	},
	8: {
		1: []*Resource{},
		2: []*Resource{{Id: 66, Img: "media/2023/06/26/06531f44140d11ee910d0242c0a80003.png", Type: 6, Term: 1, Count: 12}},
		3: []*Resource{{Id: 191, Img: "media/2022/06/20/0dc5d982f08811ec89fd0242c0a80002.png", Type: 2, Term: 1, Count: 12}},
		4: []*Resource{{Id: 256, Img: "media/2022/08/25/0fbb82be245e11edba150242c0a80002.png", Type: 2, Term: 1, Count: 12}},
	},
	9: {
		1: []*Resource{},
		2: []*Resource{{Id: 66, Img: "media/2023/06/26/06531f44140d11ee910d0242c0a80003.png", Type: 6, Term: 1, Count: 15}},
		3: []*Resource{{Id: 256, Img: "media/2022/08/25/0fbb82be245e11edba150242c0a80002.png", Type: 2, Term: 1, Count: 15}},
		4: []*Resource{{Id: 471, Img: "media/2023/06/21/c7c22ab40fe711ee83650242c0a80002.png", Type: 2, Term: 1, Count: 15}},
	},
	10: {
		1: []*Resource{{Id: 30, Img: "media/2022/08/26/6d8e20b8252b11ed83650242c0a80002.png", Type: 6, Term: 1, Count: 20}},
		2: []*Resource{{Id: 471, Img: "media/2023/06/21/c7c22ab40fe711ee83650242c0a80002.png", Type: 2, Term: 1, Count: 20}},
		3: []*Resource{{Id: 468, Img: "media/2023/06/20/f9964dd00f3911eeb05e0242c0a80002.png", Type: 2, Term: 1, Count: 20}},
		4: []*Resource{{Id: 41, Img: "media/2022/11/10/d1a187bc610211ed8dc60242c0a80002.png", Type: 4, Term: 1, Count: 20}},
	},
	11: {
		1: []*Resource{{Id: 30, Img: "media/2022/08/26/6d8e20b8252b11ed83650242c0a80002.png", Type: 6, Term: 1, Count: 20}},
		2: []*Resource{{Id: 468, Img: "media/2023/06/20/f9964dd00f3911eeb05e0242c0a80002.png", Type: 2, Term: 1, Count: 20}},
		3: []*Resource{{Id: 23, Img: "media/2023/06/27/ad6d254c14ba11ee8e180242c0a80002.png", Type: 2, Term: 1, Count: 20}},
		4: []*Resource{{Id: 75, Img: "media/2023/06/21/91639daa100911ee91480242c0a80003.png", Type: 4, Term: 1, Count: 20}},
	},
	12: {
		1: []*Resource{{Id: 67, Img: "media/2023/06/26/137521e0140d11eeb4b30242c0a80003.png", Type: 6, Term: 2, Count: 20}},
		2: []*Resource{{Id: 23, Img: "media/2023/06/27/ad6d254c14ba11ee8e180242c0a80002.png", Type: 2, Term: 2, Count: 20}},
		3: []*Resource{{Id: 406, Img: "media/2023/02/15/b56a9ebcad4711eda1420242c0a80002.png", Type: 2, Term: 2, Count: 20}},
		4: []*Resource{{Id: 74, Img: "media/2023/06/21/7f82ced0100911ee91480242c0a80003.png", Type: 4, Term: 2, Count: 20}},
	},
	13: {
		1: []*Resource{{Id: 67, Img: "media/2023/06/26/137521e0140d11eeb4b30242c0a80003.png", Type: 6, Term: 2, Count: 25}},
		2: []*Resource{{Id: 469, Img: "media/2023/06/20/3a6464dc0f3a11ee858b0242c0a80002.png", Type: 2, Term: 2, Count: 25}},
		3: []*Resource{{Id: 470, Img: "media/2023/06/20/47f558ae0f3a11eeb05e0242c0a80002.png", Type: 2, Term: 2, Count: 25}},
		4: []*Resource{{Id: 34, Img: "media/2022/09/05/802c29ec2d1711eda73f0242c0a80002.png", Type: 4, Term: 2, Count: 25}},
	},
	14: {
		1: []*Resource{{Id: 68, Img: "media/2023/06/26/1d0841c4140d11eeaccb0242c0a80003.png", Type: 6, Term: 3, Count: 25}},
		2: []*Resource{{Id: 473, Img: "media/2023/06/21/61cdf93a0fe811eea1ed0242c0a80002.png", Type: 2, Term: 3, Count: 25}},
		3: []*Resource{{Id: 472, Img: "media/2023/06/21/4e0bffe60fe811eea5a30242c0a80003.png", Type: 2, Term: 3, Count: 25}},
		4: []*Resource{{Id: 40, Img: "media/2022/11/08/61d8ee125f4411edb5f60242c0a80002.png", Type: 4, Term: 3, Count: 25}},
	},
	15: {
		1: []*Resource{{Id: 68, Img: "media/2023/06/26/1d0841c4140d11eeaccb0242c0a80003.png", Type: 6, Term: 3, Count: 30}},
		2: []*Resource{{Id: 474, Img: "media/2023/06/21/756496020fe811ee84b90242c0a80002.png", Type: 2, Term: 3, Count: 30}},
		3: []*Resource{{Id: 475, Img: "media/2023/06/21/94d436460fe811ee97080242c0a80002.png", Type: 2, Term: 3, Count: 30}},
		4: []*Resource{{Id: 76, Img: "media/2023/06/21/47689010101411ee91480242c0a80003.png", Type: 4, Term: 3, Count: 30}},
	},
}

var FamilyLevelConfigMap map[int32]FamilyTaskBoxGift

func ctxx(ctx context.Context) context.Context {
	ctx = context.WithValue(ctx, "ddd", 222)
	return ctx
}
func SAArrayToString(arr []string, split string) string {
	if len(arr) == 0 {
		return ""
	}
	arg := arr[0]
	for i := 1; i < len(arr); i++ {
		arg += split
		arg += arr[i]
	}

	return arg
}
func SAStringToArray(str string, split string) []string {
	if len(str) == 0 {
		return []string{}
	}

	return strings.Split(str, split)
}

var (
	SvipInterestsMap SvipInterestsInformationMap
)

type SvipInterestsInformationMap map[string]SvipInterestsInformation
type SvipInterestsInformation struct {
	Resources []SvipInterestsResource `json:"resources"`
}
type SvipInterestsResource struct {
	ID   int    `json:"id"`
	Test string `json:"test"`
	Type int    `json:"type"`
}

func changeSlice(s1 []int) {
	arrayAddr := (*int)(unsafe.Pointer(&s1[0]))
	fmt.Printf("切片底层数组的地址2: %p, len: %v\n", arrayAddr, len(s1))
	s1 = append(s1, 10)
	arrayAddr = (*int)(unsafe.Pointer(&s1[0]))
	fmt.Printf("切片底层数组的地址3: %p, len: %v\n", arrayAddr, len(s1))
}
<<<<<<< Updated upstream

func main() {
=======
func hash(id string, num int) int {
	h := getCrc32(id)
	h = Abs(h)
	hashValue := (int)(h / int64(num) % int64(num))
	return hashValue
}

func Abs(n int64) int64 {
	y := n >> 63
	return (n ^ y) - y
}

func getCrc32(id string) int64 {
	return int64(crc32.ChecksumIEEE([]byte(id)))
}
func DateStringYMDHMSFromTimestampMilli(now int64, tz_seconds int) string {
	localtime := now/1000 + int64(tz_seconds)
	return time.Unix(localtime, 0).In(time.UTC).Format("2006-01-02 15:04:05")
}
func MGet(t ...string) {

}

func sender(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch) // 关闭通道
}

type MyObject struct {
	Number int
}

func main() {
	{
		// 创建一个对象池
		pool := &sync.Pool{
			New: func() interface{} {
				fmt.Println("Creating new object")
				return &MyObject{}
			},
		}

		// 将对象放入池中
		//obj := &MyObject{Number: 1}
		//fmt.Printf("1 Retrieved object: %v, addr: %v, addr2: %v\n", obj.Number, &obj, obj)
		//pool.Put(obj)

		// 从池中获取对象
		retrievedObj := pool.Get().(*MyObject)
		fmt.Printf("2 Retrieved object: %v, addr: %v\n", retrievedObj.Number, &retrievedObj)

		// retrievedObj.Number = 10
		// 对象使用完后放回池中
		pool.Put(retrievedObj)

		// 再次获取对象时，会直接从池中获取
		retrievedObj2 := pool.Get().(*MyObject)
		fmt.Printf("3 Retrieved object again: %v, addr: %v\n", retrievedObj2.Number, &retrievedObj2)
	}
	{
		ch := make(chan int)
		go sender(ch)

		for value := range ch {
			fmt.Println("Received:", value)
		}

		fmt.Println("Done receiving!")
		// i := 0
		// atomic.CompareAndSwapUint32((*uint32)(i), 0, 1)
	}
	{
		const layout = "2006-01-02 15:04:05"
		dateTimeStr := "2024-05-01 00:00:00"

		// 解析日期时间字符串
		dateTime, err := time.Parse(layout, dateTimeStr)
		if err != nil {
			panic(err)
		}

		// 获取UTC时间戳（自1970年1月1日00:00:00 UTC以来的秒数）
		timestamp := dateTime.Unix()

		// 打印UTC时间戳
		fmt.Printf("UTC Timestamp: %d\n", timestamp)

		// 如果你需要纳秒级的时间戳
		nanoTimestamp := dateTime.UnixNano()
		fmt.Printf("UTC Nano Timestamp: %d\n", nanoTimestamp)

		fmt.Printf("format string: %d\n", time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC).Unix())

		fmt.Printf("format string2: %s\n", time.Unix(time.Date(2024, 5, 1, 0, 0, 0, 0, time.UTC).Unix(), 0).UTC().Format(layout))

		fmt.Printf("format string3: %s\n", time.Unix(1713810780, 0).UTC().Format(layout))
		fmt.Printf("format string4: %s", time.Now().Format(layout))

	}
	{
		slice := []int{0, 1, 2}
		fmt.Printf("orislice: %v, len: %v, cap: %v\n", slice, len(slice), cap(slice))
		slice = slice[1:]
		fmt.Printf("slice: %v, len: %v, cap: %v\n", slice, len(slice), cap(slice))

		slice2 := []int{0, 1, 2}
		fmt.Printf("orislice2: %v, len: %v, cap: %v\n", slice2, len(slice2), cap(slice2))
		fmt.Printf("slice2[:0]: %v, len: %v, cap: %v\n", slice2[:0], len(slice2[:0]), cap(slice2[:0]))
		slice2 = append(slice2[:0], slice2[1:]...)
		fmt.Printf("slice2: %v, len: %v, cap: %v\n", slice2, len(slice2), cap(slice2))
	}
	{
		s := []string{"1", "2"}
		MGet(s...)

	}
	{
		s := []string{"1", "2"}
		fmt.Println(s[len(s):])
	}
	{
		fmt.Printf("ceshi: %v", 123456-102345)
	}
	{
		visited := map[*misccode.TreeNode]bool{}
		var x *misccode.TreeNode
		visited[x] = true
		fmt.Printf("%+v", visited)
		if v, ok := visited[x]; ok && v {
			fmt.Println("9999")
		}
	}
	{
		fmt.Println(DateStringYMDHMSFromTimestampMilli(1*86400*1000+0*3600*1000, 0))
	}
	{
		now := time.Now().UnixMilli()
		fmt.Println("ttime, now: ", now, "end: ", 1706262000000)
>>>>>>> Stashed changes

	{
		testMap := map[int]*SvipInterestsResource{1: {ID: 1, Test: "1"}, 2: {ID: 2}, 3: {ID: 3}}

		newTestMap := make(map[int]*SvipInterestsResource)
		for k, v := range testMap {
			newTestMap[k] = v
		}
		fmt.Println(newTestMap)
	}
	{
		stest := make([]int, 10, 12)
		stest1 := stest[8:]
		arrayAddr := (*int)(unsafe.Pointer(&stest1[0]))
		fmt.Printf("切片底层数组的地址: %p, len: %v\n", arrayAddr, len(stest1))
		changeSlice(stest1)
		stest1 = stest1[:3]
		fmt.Printf("s: %v, len of s: %d, cap of s: %d\n", stest, len(stest), cap(stest))
		fmt.Printf("s1: %v, len of s1: %d, cap of s1: %d, stest1[2]: %v\n", stest1, len(stest1), cap(stest1), stest1[2])
		arrayAddr = (*int)(unsafe.Pointer(&stest1[0]))
		fmt.Printf("切片底层数组的地址: %p, len: %v\n", arrayAddr, len(stest1))
	}

	mm := map[int]int{1: 1}
	misccode.MapCopy(mm)
	fmt.Println(mm)

	numsRoate := []int{1, 2, 3, 4, 5, 6, 7}
	misccode.Rotate3(numsRoate, 3)
	fmt.Println(numsRoate)

	numsInsertArray := []int{1, 2, 3, 4, 5}
	fmt.Println(reflect.TypeOf(numsInsertArray).Kind())
	fmt.Println(misccode.InsertArray(numsInsertArray, 6, 100))

	fmt.Println(misccode.MaxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))

	fmt.Println(misccode.MinWindow("ADOBECODEBANC", "ABC"))
	var ans []interface{}
	// ans = []SvipInterestsResource{}
	//if vres, ok := ans.(int); ok {
	//	fmt.Println(vres)
	//}
	fmt.Println(reflect.TypeOf(ans))

	misccode.Trap42([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	misccode.FindMedianSortedArrays([]int{0, 1}, []int{2})

	mn := int64(1)
	fmt.Println(mn == 1)

	misccode.PartitionLabels("ababcbacadefegdehijhklij")
	cmf := misccode.ConstructorMedianFinder2()
	cmf.AddNum(-1)
	cmf.AddNum(-2)
	cmf.AddNum(-3)
	fmt.Println(cmf.FindMedian())
	cmf.AddNum(-4)
	cmf.AddNum(-5)
	//cngmf.AddNum(0)
	//cmf.AddNum(6)
	//cmf.AddNum(3)
	//cmf.AddNum(1)
	//cmf.AddNum(0)
	//cmf.AddNum(0)

	svipR := &SvipInterestsResource{
		ID:   1,
		Type: 10,
		Test: "98989",
	}
	svipRB, _ := json.Marshal(svipR)
	fmt.Println(string(svipRB))
	svipRB = []byte("{\"id\":1,\"test\":\"98989\",\"type\":10, \"yuyu\":78}")
	fmt.Println(string(svipRB))
	svipS := &SvipInterestsResource{}
	json.Unmarshal(svipRB, svipS)
	fmt.Println(svipS)

	// test1 := []int{10, -3, 0, 5, 9}
	head108 := misccode.SortedArrayToBST([]int{1, 3})
	fmt.Println(head108)

	newNodes := make([]*misccode.ListNode, 5)
	fmt.Println(newNodes[1].Val)

	dHead := &misccode.ListNode{}
	dh := dHead
	for i := 0; i < 2; i++ {
		dh.Next = &misccode.ListNode{Val: i + 1}
		dh = dh.Next
	}

	misccode.ReverseKGroup(dHead.Next, 2)

	fmt.Println(strings.Join([]string{"sfag"}, ","))
	fmt.Println(strings.Split("sfag", ","))

	ists := make(map[string]SvipInterestsInformation)
	for i := 1; i <= 4; i++ {
		ists[fmt.Sprint(i)] = SvipInterestsInformation{
			Resources: []SvipInterestsResource{
				{ID: 1, Type: 20},
				{ID: 2, Type: 30},
			},
		}
	}
	istsB, err := json.Marshal(ists)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(istsB))

	f1, err := os.Open("/Users/mico/svip_interests.json")
	defer f1.Close()
	if err != nil {
		fmt.Println(err)
	}

	b, err := ioutil.ReadAll(f1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
	jsonData := `{"1":"{\"resources\": [{\"id\": 40,\"type\":2},{\"id\": 20,\"type\":1}]","2":"{\"resources\": [{\"id\": 40,\"type\":2},{\"id\": 20,\"type\":1}]","3":"{\"resources\": [{\"id\": 40,\"type\":2},{\"id\": 20,\"type\":1}]","4":"{\"resources\": [{\"id\": 40,\"type\":2},{\"id\": 20,\"type\":1}]"}`

	interests := make(map[string]SvipInterestsInformation)
	err2 := json.Unmarshal([]byte(jsonData), &interests)
	if err2 != nil {
		fmt.Println(err2)
	}

	var strrr []string
	fmt.Println(SAArrayToString(strrr, "|"))
	fmt.Println(len(SAArrayToString(strrr, "|")))
	fmt.Println(len(SAArrayToString([]string{}, "|")))
	fmt.Println(len(""))

	fmt.Println(SAArrayToString([]string{"", "", "", ""}, "|"))
	fmt.Println(len(SAStringToArray("|||", "|")))

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	misccode.Rotate(nums, 3)
	fmt.Println(nums)

	fmt.Printf("%T\n", make([]int, 0))
	fmt.Printf("%T\n", make(map[int]int))
	fmt.Printf("%T\n", make(chan int))
	i := 0
	fmt.Printf("%T\n", &i)

	fmt.Println(len(strings.Split("ga", ",")))
	fmt.Println(len(strings.Split("falkjg", ",")))
	fmt.Println(len(strings.Split("fal,kjg", ",")))
	ctx := context.Background()
	ctx = context.WithValue(ctx, "ddd", 111)
	//ctx = context.WithValue(ctx, "ddd", 222)
	ctx = ctxx(ctx)
	ctxx, _ := context.WithTimeout(ctx, time.Duration(10))
	fmt.Println(ctxx.Value("ddd"))

	misccode.MaximalRectangle([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}})
	fmt.Println(int('1' - '0'))
	misccode.LongestValidParentheses("(()))())(")

	fmt.Println(3 < 3) // false
	var m map[int]Resource
	fmt.Println(m == nil)
	fmt.Println(m[1].Count == 1)

	fmt.Println(len(strings.Split("11", ",")))
	fmt.Println(strings.Split("11", ",")[0])
	coins, err := json.Marshal(FamilyBoxCoinss)
	if err == nil {
		fmt.Println(string(coins))
	}
	active, err := json.Marshal(FamilyBoxActive)
	if err == nil {
		fmt.Println(string(active))
	}
	resources, err := json.Marshal(FamilyBoxResources)
	if err == nil {
		fmt.Println(string(resources))
	}

	//fbc := FamilyBoxCoinsT{}
	//err = json.Unmarshal(coins, &fbc)
	//if err == nil {
	//	fmt.Println(fbc)
	//}

	page := []string{"fda", "gf", "qw", "ert"}
	fmt.Println(page[3:5])
	fmt.Println("2023-06-25 02:52:08" > "2023-06-21 06:32:00")

	fmt.Println(strings.Split("sfag", ","))
	fmt.Println(strings.Join([]string{"fa", "gd", "nnmm"}, ","))
	fmt.Println(strings.Split(strings.Join([]string{"fa", "gd", "nnmm"}, ","), ","))

	extend := &struct {
		GiftType     int64 `json:"giftType"`
		FromFunction int64 `json:"from_function"`
	}{}

	err = json.Unmarshal([]byte("{\"avtSessionId\":\"\",\"fromUserFamily\":2002806,\"from_function\":0,\"giftType\":3,\"giftsConsumeCoins\":3999,\"toUserFamily\":2002806}"), extend)
	fmt.Println(err)

	misccode.CalcEquation([][]string{{"a", "b"}, {"c", "d"}}, []float64{1.0, 1.0}, [][]string{{"a", "c"}, {"b", "d"}, {"b", "a"}, {"d", "c"}})

	fmt.Println(fmt.Sprintf("falkjfa ", 111))
	fmt.Println(FamilyLevelConfigMap[1].GiftAging)
	//[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]]
	//[[1,4,7,11,15],[2,5,8,12,19],[3,6,9,16,22],[10,13,14,17,24],[18,21,23,26,30]] 20
	// [[1,2,3,4,5],[6,7,8,9,10],[11,12,13,14,15],[16,17,18,19,20],[21,22,23,24,25]] 19
	fmt.Println(misccode.SearchMatrix([][]int{{1, 2, 3, 4, 5}, {6, 7, 8, 9, 10}, {11, 12, 13, 14, 15}, {16, 17, 18, 19, 20}, {21, 22, 23, 24, 25}}, 19))

	ctx = context.Background()
	ctx = context.WithValue(ctx, "TRACEID", 1234567890)

	misccode.DEBUGT(nil, "TEST DEBUGT: %s", "test debugt")

	ftbg := []FamilyTaskBoxGift{
		{
			GiftImg:   "giftImg1",
			GiftName:  "giftName1",
			GiftAging: "giftAging1",
			GiftCount: 14,
		},
		{
			GiftImg:   "giftImg2",
			GiftName:  "giftName2",
			GiftAging: "giftAging2",
			GiftCount: 24,
		},
	}
	ftbgB, _ := json.Marshal(&ftbg)
	fmt.Println(string(ftbgB))

	ftbg1 := []FamilyTaskBoxGift{}
	json.Unmarshal(ftbgB, &ftbg1)
	fmt.Println(ftbg1)

	fmt.Println(time.Now().Sub(time.Unix(1685607961, 0)).Hours() < 24)

	fmt.Println(time.Now().Format("2006-01-02"))

	fmt.Println(time.Unix(0, 1665705599000).Format("2006-01-02 15:04:05"))

	xt := []byte{8,
		242,
		247,
		54,
		18,
		38,
		105,
		109,
		97,
		103,
		101,
		47,
		53,
		54,
		51,
		99,
		99,
		53,
		48,
		57,
		101,
		100,
		55,
		98,
		54,
		101,
		48,
		55,
		55,
		54,
		57,
		100,
		54,
		51,
		55,
		99,
		48,
		101,
		56,
		101,
		57,
		52,
		53,
		54,
		26,
		30,
		229,
		147,
		136,
		229,
		147,
		136,
		229,
		147,
		136,
		229,
		147,
		136,
		229,
		147,
		136,
		229,
		147,
		136,
		229,
		147,
		136,
		229,
		147,
		136,
		229,
		147,
		136,
		229,
		147,
		136,
		34,
		2,
		69,
		71,
		96,
		40}
	fmt.Println(string(xt))
	misccode.ThreeSumClosest([]int{-1, 2, 1, -4}, 1)
	misccode.SearchX([]int{1, 2, 3, 5, 4, 2})
	fmt.Println(test2())
	fmt.Println(test())
	tTreeNodes := []*misccode.TreeNode{}
	var ttn *misccode.TreeNode
	tTreeNodes = append(tTreeNodes, ttn)
	fmt.Println(len(tTreeNodes[1:1]))
	fmt.Println(len(tTreeNodes[1:]))

	rt := &misccode.TreeNode{Val: 1}
	rt.Left = &misccode.TreeNode{Val: 2}
	rt.Right = &misccode.TreeNode{Val: 5}
	rt.Left.Left = &misccode.TreeNode{Val: 3}
	rt.Left.Right = &misccode.TreeNode{Val: 4}
	rt.Right.Right = &misccode.TreeNode{Val: 6}
	misccode.Flatten(rt)
	misccode.Insert()
	misccode.SearchInsert([]int{1, 2, 3, 4, 5}, 6)
	fmt.Println(misccode.MaxSum([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}))

	misccode.MinSubArrayLen3(7, []int{2, 3, 1, 2, 4, 3})
	misccode.MinSubArrayLen2(15, []int{1, 2, 3, 4, 5})
	misccode.FindAnagrams("ababababab", "aab")
	misccode.ShortestSubarray([]int{17, 85, 93, -45, -21}, 150)
	misccode.MinSubArrayLen(7, []int{2, 3, 1, 2, 4, 3})
	misccode.NumberOfSubarrays([]int{2, 2, 2, 1, 2, 2, 1, 2, 2, 2}, 2)
	misccode.Subsets([]int{1, 2, 3})
	misccode.SliceT()
	misccode.ReconstructQueue([][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}})
	misccode.DecodeString("3[a]2[bc]")

	fmt.Println(misccode.MaxValue([]int{1, 3, 4}, []int{15, 20, 30}, 4))
	misccode.MaximalSquare([][]byte{{'1', '0', '1', '0', '0'}, {'1', '0', '1', '1', '1'}, {'1', '1', '1', '1', '1'}, {'1', '0', '0', '1', '0'}})
	misccode.MaxProduct([]int{2, 3, -2, 4})
	misccode.Merge([][]int{{2, 3}, {2, 2}, {3, 3}, {1, 3}, {5, 7}, {2, 2}, {4, 6}})

	misccode.LengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	misccode.NumIslands([][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}})
	misccode.CanPartitionKSubsets14([]int{4, 3, 2, 3, 5, 2, 1}, 4)
	misccode.Exist([][]byte{{'a', 'b', 'c'}, {'a', 'e', 'd'}, {'a', 'f', 'g'}}, "abcdefg")
	stt := misccode.LongestDupSubstring("")
	fmt.Println(stt)
	return
	misccode.SubsetsWithDup([]int{1, 2, 2})
	//stt := misccode.LongestDupSubstring("")
	//fmt.Println(stt)
	//return
	misccode.CanPartitionKSubsets([]int{1, 2, 2, 2, 2}, 3)

	// misccode.LongestConsecutive([]int{100, 4, 200, 1, 3, 2})
	s := "abcd"
	wordDict := []string{"a", "b", "c", "ab", "bc"}
	fmt.Println(misccode.WordBreak(s, wordDict))
	return

	s = "abc"
	enumcode.AllSubsquences(s)
	return

	s = "abc"
	enumcode.Permutations(s)
	return
	//s1 := "abc"
	//enumcode.AllSubsquences(s1)
	fmt.Println(misccode.GenerateParenthesis(3))
	fmt.Println(misccode.IsValid("()[]{}"))
	misccode.LetterCombinations2("23")
	queuecode.LevelOrder(nil)

	misccode.LengthOfLongestSubstring("abcabcbb")

	misccode.ReversePairs([]int{1, 3, 2, 3, 1})

	nums = []int{1, 4, 2, 6, 3, 0, 9}
	misccode.QuickSort2(nums)
	fmt.Println(nums)

	// n皇后
	enumcode.Nqueen(4)
	return

	// 三数之和
	// fmt.Println(misccode.ThreeNumsSum([]int{-1, 0, 1, 2, -1, -4}, 0))

	// 字符串相同字符统计,空间复杂度O(1)
	s = "abbccc"
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

	nums = []int{30, 1, 20, 4, 6}
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
