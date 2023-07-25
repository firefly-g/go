package main

/*
 go中没有错误处理机制，一般使用panic/recover模式来处理
 panic可以在任何地方引发，但recover只有在defer调用的函数中有效
*/
import (
	"fmt"
)

func main() {
	funcA()
	funcB()
	funcC()
}
func funcA() {
	fmt.Println("A")
}
func funcB() {
	//defer语句一定要在可能引发panic的语句之前使用
	defer func() {
		err := recover()
		fmt.Println(err)
		fmt.Println("释放其他因为error而未关闭的资源")
	}()

	panic("Error!")
	fmt.Println("B")
}
func funcC() {
	fmt.Println("C")
}
