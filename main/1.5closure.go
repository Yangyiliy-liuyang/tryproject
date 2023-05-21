package main

func main() {
	r1 := increment()
	println(r1)

	v1 := r1()
	println(v1)
	v2 := r1()
	println(v2)
	println(r1())
	println(r1())
	println(r1())

	r2 := increment()
	println(r2)
	v3 := r2()
	println(v3)
	println(r1())
	println(r1())
	println(r1())

}

// 自增
func increment() func() int {
	i := 0
	fun := func() int {
		i++
		return i
	}
	return fun
}
