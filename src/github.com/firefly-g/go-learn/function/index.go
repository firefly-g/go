package main

import (
	"fmt"
	"log"
)

func main() {
	//使用defer延迟运行语句 多用于在函数结束之前释放资源*（文件句柄、数据库连接、socket连接）
	toDeferTest()
	//经典应用
	//toDefetFunc()
	//作用域
	//toContext()
	/*
		实现一个斐波那契数列
	*/
	//toRecursive()
	/*
		使用defer延迟返回,用于在return之后修改返回的数据使用
	*/
	//toDefer()
	/*
		使用闭包 将函数作为返回值
	*/
	//toClosure()

}
func toDeferTest() {
	fmt.Println("start")
	defer fmt.Println("我是一")
	defer fmt.Println("我是二")
	defer fmt.Println("我是三")
	fmt.Println("end")
}
func toDefetFunc() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}
func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
} //5
func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
} //6

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x //c此时等于 y = x = 5
} //5
func f4() (x int) {
	defer func(x int) {
		x++
	}(x) //将x传入 此时为值传递 不会改变原来的变量
	return 5
} //5

var x = 100

func toContext() {
	print1()
}
func print1() {
	x := 200
	fmt.Println(x)
}

func toRecursive() {
	var result []int
	for i := 0; i <= 10; i++ {
		temp := fibonacci(i)
		result = append(result, temp)
	}
	log.Print(result)
}
func fibonacci(i int) (ret int) {
	if i <= 1 {
		ret = 1
	} else {
		ret = fibonacci(i-1) + fibonacci(i-2)
	}
	return
}
func toDefer() {
	log.Print(f())
}
func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}

func toClosure() {
	f1 := add1()
	f2 := add2(2)
	log.Print(f1(2))
	log.Print(f2(3))
	/**
	另一种方式实现
	*/
	f3 := add3()
	log.Print(f3(2))
	log.Print(f3(10))
	log.Print(f3(100))
	/*
		使用闭包实现斐波那契数列
	*/
	toFibo()

}
func add1() func(a int) int {
	return func(a int) int {
		return a + 2
	}
}
func add2(x int) func(b int) int {
	return func(b int) int {
		return x + b
	}
}
func add3() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x

	}

}
func toFibo() {
	f := fibo()
	for i := 0; i < 10; i++ {
		log.Print(f())
	}
}
func fibo() func() int {
	pre, cur := 0, 1
	return func() int {
		item := pre
		pre, cur = cur, pre+cur
		return item
	}
}
