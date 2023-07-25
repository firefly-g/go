package main

import "log"

func main() {
	//初始化切片
	//toSlice()

	//将切片传递给函数
	//toParams()

	//用make创建一个切片
	//toCreateSlice()

	//使用make()和new()的区别
	//toDifference()

	//切片重组
	//toReslice()

	//切片的复制于追加
	to_copy_append_slice()

}
func toSlice() {
	var arr1 [6]int
	slice1 := arr1[2:5]
	//初始化
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}
	//打印切片
	for i := 0; i < len(slice1); i++ {
		log.Print(i, slice1[i])
	}
	//切片容量
	log.Print("切片容量:", cap(slice1))
	//切片再切片
	slice2 := slice1[:]
	log.Print("切片slice2:", slice2)
	//切片扩容
	slice1 = slice1[0:4]
	log.Print("更新后的slice1:", slice1) //[2 3 4 5]
	toTest()
}
func toTest() {
	//给定一个切片
	slice := []byte{'g', 'o', 'l', 'a', 'n', 'd'}
	//切片
	log.Print(slice[:])  //[103 111 108 97 110 100]
	log.Print(slice[2:]) // [108 97 110 100]
	log.Print(slice[:len(slice)])
}

func toParams() {
	arr := []int{1, 2, 3, 4}
	slice := arr[:]
	//将数组通过切片传递
	sum := sum(slice)
	log.Print(sum)
}
func sum(x []int) int {
	var sum int
	for _, item := range x {
		sum += item
	}
	return sum
}

func toCreateSlice() {
	//当相关函数还没被定义时，可以使用make()创建一个切片，同时创建好相关数组
	var slice []int = make([]int, 5)      //5为数组的长度也是slice的初始长度
	var slice1 []int = make([]int, 5, 10) //10为数组的长度  5为slice的初始长度
	log.Print(slice, cap(slice))
	log.Print(slice1, cap(slice1))

}

func toDifference() {
	/*
		make和new函数都是在堆上分配内存，但是他们的行为不同，使用不同的类型

		new(T)为每个新的类型T分配一片内存，初始化为0并返回类型为*T的内存地址 相当于&T{}
		new适用于值类型 如数组和结构体

		make(T) 返回一个类型为T的初始值，它只适用于3中内建的引用类型:切片、map、channel

		换言之， new函数分配内存 make函数初始化

		**/
	//使用new
	var p1 *[]int = new([]int) // *p == nil; with len and cap 0
	//p2:=new([]int)
	print(p1)

	//使用new
	var n []int = make([]int, 0) //切片已经被初始化 只是指向一个空数组
	print(n)
}

func toReslice() {
	//将切片扩容到占据整个相关数组
	slice := make([]int, 0, 10)
	for i := 0; i < cap(slice); i++ {
		slice = slice[0 : i+1]
		slice[i] = i
		log.Print("length of slice:", len(slice))
	}
	log.Print(slice[10:10])     //[]
	log.Print(slice[10 : 10+1]) //error slice bounds out of range
}

func to_copy_append_slice() {
	/*
		若想增加切片的容量，必须创建一个更大的切片并把原来的内容都拷贝过来
	*/
	slFrom := []int{1, 2, 3}
	slTo := make([]int, 10)

	//拷贝
	n := copy(slTo, slFrom)
	log.Print(n)    //拷贝的元素个数:3
	log.Print(slTo) //[1 2 3 0 0 0 0 0 0 0]

	//扩容
	slTo = append(slTo, 4, 5, 6)
	log.Print(slTo) //[1 2 3 0 0 0 0 0 0 0 4 5 6]

}
