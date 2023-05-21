package main

import "fmt"

func main() {
	var name string
	var age int8
	fmt.Scanln(&name)
	//fmt.Scanf("%s", &name)
	fmt.Scanf("%d", &age)
	//fmt.Scanf("%s %d", &name, &age)  对于scanf，这句话等价于上面两句话
	fmt.Println(name, "  ", age)
	fmt.Printf("%c", name[1])
}
