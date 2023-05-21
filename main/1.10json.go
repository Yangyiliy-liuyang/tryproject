package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	ID   int
	Name string
}

func newStudent(name string, id int) Student {
	return Student{
		ID:   id,
		Name: name,
	}
}

type Class struct {
	Title    string
	Students []Student
}

func main() {
	c1 := Class{
		Title:    "火箭少女101",
		Students: make([]Student, 0, 20),
	}
	for i := 0; i < 10; i++ {
		temstu := newStudent(fmt.Sprintf("stu%02d", i), i)
		c1.Students = append(c1.Students, temstu)
	}
	fmt.Printf("%#v", c1)
	data, err := json.Marshal(c1)
	if err != nil {
		fmt.Println("json marshal failed,err:")
		return
	}
	println()
	fmt.Printf("%T\n", data)
	fmt.Printf("%s\n", data)
	//JSON反序列化：JSON格式的字符串-->结构体
	str := `{"Title":"101","Students":[{"ID":0,"Gender":"男","Name":"张三"},{"ID":1,"Gender":"男","Name":"李白"},{"ID":2,"Gender":"男","Name":"stu02"},{"ID":3,"Gender":"男","Name":"stu03"},{"ID":4,"Gender":"男","Name":"stu04"},{"ID":5,"Gender":"男","Name":"stu05"},{"ID":6,"Gender":"男","Name":"stu06"},{"ID":7,"Gender":"男","Name":"stu07"},{"ID":8,"Gender":"男","Name":"stu08"},{"ID":9,"Gender":"男","Name":"stu09"}]}`
	c2 := &Class{}
	err = json.Unmarshal([]byte(str), c2)
	if err != nil {
		fmt.Println("json unmarshal failed!")
		return
	}
	fmt.Printf("%#v\n", c2)
}
