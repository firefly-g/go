package main

import (
	"fmt"
	"log"
)

func main() {
	//初始化map
	//toMap()
	//测试键值对是否存在
	toIsHasKey()
	//创建map类型的切片
	toSliceOfMap()
}

func toMap() {
	//声名map
	var m map[string]string
	//初始化1
	m = map[string]string{
		"name": "孙悟空",
		"age":  "24",
	}
	m["name"] = "猪八戒"
	log.Print(m)

	//初始化2 使用make()初始化
	m2 := make(map[string]int, 2)
	m2["one"] = 1
	m2["two"] = 2
	m2["three"] = 3
	log.Print(m2)

	//map是引用类型
	m3 := m2
	m3["four"] = 4
	m3["three"] = 33
	log.Print(m3)
	log.Print(m2)
}

func toIsHasKey() {
	//获取一个不存在的key是，map会返回值类型的空值
	p := map[string]int{"one": 1, "two": 2}
	val, isPresent := p["other"]
	log.Print(val, isPresent) // 0 false
	//判断某个key是否存在而不关心值时
	if _, ok := p["nothing"]; ok {
		log.Print("当前key在map中存在")
	} else {
		log.Print("当前key在map中不存在")
	}

	capitals := map[string]string{"France": "Paris", "Italy": "Rome", "Japan": "Tokyo"}
	for key := range capitals {
		fmt.Println("Map item: Capital of", key, "is", capitals[key])
	}

}

func toSliceOfMap() {
	/**
	1. 使用make()分配切片内存
	2. 使用make()分配切片中的每个map元素
	*/
	slice := make([]map[string]int, 3)
	//遍历切片，为每个map元素创建内存
	for i := range slice {
		slice[i] = make(map[string]int, 2)
		slice[i]["one"] = 1
		slice[i]["two"] = 2
	}
	log.Print("slice:", slice)
}
