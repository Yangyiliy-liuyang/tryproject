package main

import "fmt"

func f1(ch chan<- int) {
	for i := 0; i < 100; i++ {
		ch <- i
	}
	close(ch)
}
func f2(ch1 chan int, ch2 chan<- int) {
	for true {
		value, ok := <-ch1
		if !ok {
			break
		}
		ch2 <- value * value
	}
	close(ch2)
}
func main() {
	ch := make(chan int, 1) //带缓冲区通道
	ch <- 10
	x := <-ch
	fmt.Println(x)
	close(ch)

	ch1 := make(chan int, 100)
	ch2 := make(chan int, 100)

	f1(ch1)
	f2(ch1, ch2)
	for ret := range ch2 {
		fmt.Println(ret)
	}
}
