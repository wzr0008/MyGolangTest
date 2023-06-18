package Concurrency

import (
	"fmt"
	"time"
)

func sendData1(ch chan<- int) {
	ch <- 1
}
func getData1(ch <-chan int) {
	val := <-ch
	fmt.Println(val)
}
func GoroutineTest4() {
	ch := make(chan int)
	go sendData1(ch)
	go getData1(ch)
	time.Sleep(1 * time.Second)

}
