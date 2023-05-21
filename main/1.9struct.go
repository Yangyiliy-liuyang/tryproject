package main

import (
	"fmt"
)

type person struct {
	name, sex, city string
	age             int8
}

func main() {
	var p person
	p.name = "娜扎"
	p.sex = "女"
	p.city = "武汉"
	p.age = 18
	fmt.Println(p)

	//匿名结构体
	var user struct {
		Name string
		Age  int
	}
	user.Name = "小王子"
	user.Age = 18
	fmt.Printf("%#v\n", user)

	var admin = new(person)
	(*admin).name = "张三"
	admin.sex = "男"
	(*admin).city = "上海"
	fmt.Printf("%#v\n", admin)
}
