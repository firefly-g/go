package main

import (
	"fmt"
	"log"
)

/*
go中的接口非常灵活 它可以实现很多面向对象的特性，它提供了一种方式来说明对象的行为
*/
type Shaper interface {
	Area() float32
}
type Square struct {
	side float32
}
type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}
func (s *Square) Area() float32 {
	return s.side * s.side
}
func main() {
	//sq1 := Square{4}
	sq1 := new(Square)
	sq1.side = 5
	var shaperInt Shaper
	shaperInt = sq1
	fmt.Printf("The square has area: %f\n", shaperInt.Area())
	//接口实现多态
	toInterfacesPoly()
	//接口嵌套接口
	toNestedInterface()
	//空接口
	toEmptyInterface()

}

func toInterfacesPoly() {
	s := &Square{5}
	q := Rectangle{10, 3}
	//将结构变量存入数组中 赋值给接口变量
	shapes := []Shaper{s, q}
	//遍历数组 一次执行方法
	for i, _ := range shapes {
		log.Print("shape detail:", shapes[i])
		log.Print("Area of this shape is:", shapes[i].Area())
	}

}

type ReadWrite interface {
	Read(b Buffer) bool
	Write(b Buffer) bool
}
type Lock interface {
	Lock()
	UnLock()
}
type File interface {
	ReadWrite
	Lock
	Close()
}
type Buffer struct {
	name string
	size float32
}

func (b Buffer) Read(br Buffer) bool {
	log.Print("读文件中")
	return true
}
func (b Buffer) Write(br Buffer) bool {
	log.Print("取文件中")
	return true
}
func (b Buffer) Close() {
	log.Print("关闭buffer")
}
func (b Buffer) Lock() {
	log.Print("上锁")
}
func (b Buffer) UnLock() {
	log.Print("解锁")
}

func toNestedInterface() {
	/***
	一个接口可以包含一个或多个其他的接口，这相当于直接将这些内嵌接口的方法列举在外层接口中一样。
	**/
	f := Buffer{"文件", 1024}
	var fileIntf File
	fileIntf = f
	log.Print(fileIntf.Write(f))

}

func toEmptyInterface() {

	/*
	   所有的类型都实现了空接口 也就是任意类型的变量都能保存到空接口 通常写为interface{}
	   **/
	var m1 map[string]interface{}
	m1 = make(map[string]interface{}, 12)
	m1["name"] = "gt"
	m1["age"] = 1000
	m1["hobby"] = [...]string{"swimming", "reading"}

	//空接口作为函数的参数
	show(m1)
	show(true)
	//那么 如何判断空接口中值的类型呢？ 使用断言！ 语法 x.(T)
	//x 表示类型为interface()的变量 T表示断言x可能的类型
	var m2 interface{}
	m2 = "hello"
	v, ok := m2.(string)
	fmt.Println(v)
	fmt.Println(ok)

}
func show(a interface{}) {
	fmt.Println(a)
}
