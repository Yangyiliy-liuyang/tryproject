package main

import "fmt"

func main() {
	f1()
	f2 := f1
	f2()
	f3 := func() {
		fmt.Println("func")
	}
	f3()
	func() {
		fmt.Println("func")
	}()
	func(a, b int) {
		fmt.Println("func", a, b)
	}(1, 3)
	r := func(a, b int) int {
		fmt.Println("func", a, b)
		return a + b
	}(1, 3)
	print(r)
}

// f1
func f1() {
	fmt.Println("func1")
}
