package main

import "fmt"

func main() {
	var array = [...]int{1: 22, 4: 212, 6: 2}
	fmt.Printf("%v", array)
	for index, value := range array {
		println(index, value)
	}

	a := [...][2]string{
		{"北京", "上海"},
		{"广州", "深圳"},
		{"成都", "重庆"},
	}
	fmt.Println(a)       //[[北京 上海] [广州 深圳] [成都 重庆]]
	fmt.Println(a[2][1]) //支持索引取值:重庆

	for _, v1 := range a {
		//fmt.Println(v1)
		for _, v2 := range v1 {
			fmt.Printf("%s\t", v2)
		}
		fmt.Println()
	}
}
