package main

import "fmt"

var (
	name string
	age  int
	isOk bool
)

func main() {
	name = "gantian"
	age = 18
	isOk = true
	path := "\"D:\\Code\\go\"" //字符串的转义
	//多行字符串
	str := `
	归去
	也无风雨
	也无晴
	`
	// fmt.Println("hello world!")
	// fmt.Println("name is:", name)
	fmt.Println("path is:", path)
	fmt.Println("str is:", str)

	//循环
	for _, item := range str {
		fmt.Printf("%c\n", item)
	}
	//跳出循环
	for i := 0; i < 6; i++ {
		if i == 5 {
			break
		}
	}
	//跳出此次循环
	for i := 0; i < 5; i++ {
		if i == 5 {
			continue
		}
		fmt.Println(i)
	}
	//类型转换
	s2 := "小白兔"
	s3 := []rune(s2) //转换为一个rune切片
	s3[0] = '大'
	fmt.Println(string(s3)) //rune切片转换为字符串

	n := 10
	var f float64
	f = float64(n)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
