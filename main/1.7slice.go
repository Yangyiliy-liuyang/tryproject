package main

import (
	"fmt"
	"sort"
)

func main() {
	//请使用内置的sort包对数组var a = [...]int{3, 7, 8, 9, 1}进行排序
	var a = [...]int{3, 7, 8, 9, 1}
	sort.Ints(a[:])
	fmt.Println(a)
	var b []int         //切片声明
	b = []int{12, 2, 3} //初始化
	println(b)
	fmt.Printf("%v", b)
}
