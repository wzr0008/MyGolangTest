package Concurrency

import "fmt"

func produce(a, b int) chan int {
	x := make(chan int)
	go func() {
		x <- a + b
	}()
	return x
}
func consumer(ch chan int) {
	go func() {
		for {
			a := <-ch
			fmt.Println(a)
		}
	}()
}
func GoroutineTest3() {
	consumer(produce(3, 4))
}
