package main

import "fmt"

func main() {
	//if语句
	// toIf()

	//switch 语句
	// toSwitch()

	//数组
	// toArray()

	//切片
	// toSlice()

	//切片扩容
	// toSliceAppend()

	//复制切片
	// toSliceCopy()

	//从切片中删除元素
	// toSliceDelete()

	//指针
	// toPointer()

	//map
	toMap()

	//元素类型为map的切片
	toSliceOfMap()
	//元素类型为切片的map
	toMapOfSlice()
}

func toIf() {
	age := 18
	if age > 18 {
		fmt.Println("已成年")
	} else {
		fmt.Println("未成年")
	}
}
func toSwitch() {
	num := 2
	switch num {
	case 0:
		fmt.Println("0")
	case 1:
		fmt.Println("1")
	case 2, 3, 4, 5:
		fmt.Println("2")
	default:
		fmt.Println("未找到")
	}
}
func toArray() {
	var arr [3]int
	//初始化
	arr = [3]int{1, 2, 3}
	//自动推断数组长度
	arr1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 8}
	//指定索引来初始化
	arr2 := [3]int{0: 0, 2: 2}

	fmt.Println(arr)
	fmt.Println(arr1)
	fmt.Println(arr2)

	//多维数组
	var arr3 [3][2]int
	arr3 = [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}
	fmt.Println(arr3)

	//数组是值类型
	b1 := [3]int{1, 2, 3}
	b2 := b1
	b2[0] = 100
	fmt.Println(b2)
	fmt.Println(b1)

	//数组的遍历
	toArrayForEach()
}
func toArrayForEach() {
	citys := [...]string{"北京", "上海", "广东", "湖北"}
	arr3 := [3][2]int{
		[2]int{1, 2},
		[2]int{3, 4},
		[2]int{5, 6},
	}

	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	//多维数组的遍历
	for _, item := range arr3 {
		for _, item1 := range item {
			fmt.Println(item1)
		}

	}
}
func toSlice() {
	///切片的定义
	var s []int     //整行切片切片
	var s2 []string //字符串切片
	var b = []bool{true, false}
	fmt.Println(s, s2, b)
	fmt.Println(s == nil) //nil 表示为空

	//切片的长度和容量
	fmt.Println(len(s2), cap(b))

	//由已知数组得到切片
	a1 := [...]int{1, 2, 3}
	a2 := a1[0:1]
	a3 := a1[0:]
	a4 := a1[:3]
	a5 := a1[:]

	fmt.Println(cap(a2), a3, a4, a5)

	//切片再切片
	b1 := [...]int{1, 2, 3, 3, 4, 5, 6, 6}
	b2 := b1[1:4]
	fmt.Println(cap(b2)) //7
	b3 := b2[1:3]
	fmt.Println(cap(b3)) //6

	//切片是一个引用类型 ，指向底层的一个数组
	//make函数创建切片
	c1 := make([]int, 5, 10)
	fmt.Println(c1, cap(c1))
}
func toSliceAppend() {
	s1 := []string{"东", "西", "南", "北"}
	//扩容 必须用原来的变量接收
	fmt.Println(cap(s1))
	s1 = append(s1, "方向", "hahah")
	fmt.Println(s1)
	//扩容后 容量发生变化
	fmt.Println(cap(s1))
	s2 := []string{"武汉", "西南", "重庆"}
	s1 = append(s1, s2...)
	fmt.Println(cap(s1))
}
func toSliceCopy() {
	a1 := []int{1, 2, 3, 4}
	a2 := a1
	var a3 = make([]int, 3, 3)
	copy(a3, a2)
	a1 = append(a1, 5)
	// a1[5] = 100
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}
func toSliceDelete() {
	a1 := []int{1, 2, 34, 4, 5, 5, 6, 6, 6}
	//删除数组中索引为2的元素
	a1 = append(a1[:2], a1[3:]...)
	fmt.Println(a1)
	a2 := a1[:]
	a2 = append(a2[:3], a2[4:]...)
	fmt.Println(a2)
	fmt.Println(cap(a1))

	s1 := [...]int{1, 3, 5}
	s2 := s1[:]
	s2 = append(s2[:1], s2[2:]...)
	fmt.Println(s2) //[1 5]
	fmt.Println(s1) //[1 5 5]
}
func toPointer() {
	//& 取地址    *根据地址取值
	n := 10
	address := &n
	fmt.Println(address)
	fmt.Printf("%T\n", address) //*int :表示int类型的指针
	fmt.Println(*address)       //10

	//new 函数申请一个内存地址
	var a *int
	a2 := new(int)
	fmt.Println(a)
	fmt.Println(a2, *a2)
}
func toMap() {
	/**
	map类型变量默认初始值为nil 需要使用make()函数来分配内存
	*/
	var m1 map[string]int
	m1 = make(map[string]int, 2) //应确定好map容量 避免在程序运行期间动态扩容
	m1["甘甜"] = 25
	m1["甘露"] = 20
	fmt.Println(m1)
	fmt.Println(m1["甘露"])
	value, ok := m1["ok"]
	if !ok {
		fmt.Println("不存在ok")
	} else {
		fmt.Println(value)
	}
	/*
		map的遍历
	*/
	for key, value := range m1 {
		fmt.Println(key, value)
	}
	//只打印key
	for key := range m1 {
		fmt.Println(key)
	}
	//只打印value
	for _, value := range m1 {
		fmt.Println(value)
	}

	//map的删除
	delete(m1, "甘甜")
	fmt.Println(m1)
}
func toSliceOfMap() {
	var s1 = make([]map[string]int, 1, 10)
	//对切片中的每一个map类型数据做初始化
	s1[0] = make(map[string]int, 2)
	s1[0]["gt"] = 25
	s1[0]["gll"] = 20
	fmt.Println(s1)
}

func toMapOfSlice() {
	m1 := make(map[string][]int, 2)
	m1["gt"] = []int{1, 2, 3}
	fmt.Println(m1)

}
