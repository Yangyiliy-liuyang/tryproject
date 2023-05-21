package main

import "fmt"

func main() {
	//fmt.Printf("%T", f1)  //func()   f1不带括号，函数就是一个变量
	//fmt.Printf("%T", f1) //func(int, int)
	//fmt.Printf("%T", f1) //func(int, int) int

	var f5 func(int, int) int //定义一个函数类型的变量
	f5 = f1
	a := f5(2, 3)
	fmt.Println(a)
}
func f1(a, b int) int {
	return 1
}
