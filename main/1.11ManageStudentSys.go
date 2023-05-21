package main

import "fmt"

/*
Write a student information management system using "object oriented" thinking mode.
Students have id, name, age, score and so on
The program provides display student list, add students, edit student information, delete students and other functions.

使用“面向对象”的思维方式编写一个学生信息管理系统。
学生有id、姓名、年龄、分数等信息
程序提供展示学生列表、添加学生、编辑学生信息、删除学生等功能。
*/

// Student 学生结构体
type Student struct {
	ID    int
	Name  string
	Age   int
	Score float64
}

// StudentList 学生列表
type StudentList struct {
	students []*Student
}

// NewStudentList 创建学生列表
func NewStudentList() *StudentList {
	return &StudentList{
		students: make([]*Student, 0),
	}
}

// AddStudent  添加学生
func (sl *StudentList) AddStudent(s *Student) {
	sl.students = append(sl.students, s)
}

// ShowStudents 显示所有学生
func (sl *StudentList) ShowStudents() {
	for _, s := range sl.students {
		fmt.Printf("ID:%d Name:%s Age:%d Score:%.2f\n", s.ID, s.Name, s.Age, s.Score)
	}
}

// EditStudent  编辑学生
func (sl *StudentList) EditStudent(id int, name string, age int, score float64) {
	for _, s := range sl.students {
		if s.ID == id {
			s.Name = name
			s.Age = age
			s.Score = score
		}
	}
}

// RemoveStudent 删除学生
func (sl *StudentList) RemoveStudent(id int) {
	for i, s := range sl.students {
		if s.ID == id {
			sl.students = append(sl.students[:i], sl.students[i+1:]...)
		}
	}
}

func main() {
	sl := NewStudentList()

	// 添加学生
	sl.AddStudent(&Student{1, "张三", 30, 85})
	sl.AddStudent(&Student{2, "李四", 40, 95})

	// 显示学生
	sl.ShowStudents()

	// 编辑学生
	sl.EditStudent(1, "王五", 50, 75)

	// 删除学生
	sl.RemoveStudent(2)

	// 再显示
	sl.ShowStudents()
}
func menu() {
	fmt.Println("------ManageStudentSystem------")
	fmt.Println("-- 1.show details of students--")
	fmt.Println("-- 2.add   student           --")
	fmt.Println("-- 3.edit  student           --")
	fmt.Println("-- 4.drop  student           --")
	fmt.Println("-------------------------------")
}
