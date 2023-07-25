package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	//前缀和后缀 包含 字串位置 等
	toFilter()
}
func toFilter() {
	s1 := string("http:http//www.baidu.com")
	//前缀
	isHasPrefix := strings.HasPrefix(s1, "http")
	fmt.Print(isHasPrefix)
	//后缀
	isHasSuffix := strings.HasSuffix(s1, "com")
	fmt.Print(isHasSuffix)
	//字符串包含关系
	isContain := strings.Contains(s1, "baidu")
	fmt.Println(isContain)
	//字串在父串中首次出现的位置
	strIndex := strings.Index(s1, "www")
	log.Printf("%d", strIndex)
	//最后一次出现的位置
	strLastIndex := strings.LastIndex(s1, "www")
	log.Printf("%d", strLastIndex)
	//字符串替换
	s2 := strings.Replace(s1, "http", "https", 2)
	log.Printf("%s", s2)
	//字符串出现次数
	num := strings.Count(s1, "http")
	log.Printf("%d", num)
	//重复字符串
	s3 := strings.Repeat(s1, 2)
	log.Printf("%s", s3)
	//修改字符串大小写
	s4 := strings.ToUpper(s1)
	log.Printf("%s", s4)
	s5 := strings.ToLower(s1)
	log.Printf("%s", s5)
	//删除某字符 开头和结尾的字符串去掉
	s6 := strings.Trim(s1, "com")
	log.Printf("%s", s6)
	//分割字符串 返回一个slice
	//利用空白格进行分割
	s7 := strings.Fields(s1)
	log.Printf("%s", s7)
	//自定义符号分割
	s8 := strings.Split(s1, ".")
	log.Printf("%s", s8)
	//将slice拼接到字符串
	s7 = append(s7, "google", "firefox")
	s9 := strings.Join(s7, "-")
	log.Printf("%s", s9)

}
