package main

import (
	"fmt"
	"unicode"
)

/*
	函数
	func 函数名 (参数)（返回值）{
		函数体
	}
	go中没有默认参数这个概念

*
*/
func main() {
	//可变长参数
	toVariParams("1", 2, 3, 4, 5, 6)
	//判断字符串中汉字的数量
	toFilterStr()
	//判断回文
	toHuiwen()

}
func toVariParams(x string, y ...int) {
	fmt.Println(y)
}
func toFilterStr() {
	s1 := "我的名字叫gantian"
	count := 0
	for _, item := range s1 {
		//判断是否为汉字
		if unicode.Is(unicode.Han, item) {
			count++
		}
	}
	fmt.Println(count)
}
func toHuiwen() {
	s1 := "哈abcdedcba哈"
	for i, _ := range s1 {
		if i < len(s1)/2 && s1[i] != s1[len(s1)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")
}
