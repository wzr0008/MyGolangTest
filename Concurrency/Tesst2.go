package Concurrency

import (
	"fmt"
	"time"
)

func sendData(ch chan string) {
	defer close(ch)
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"

}
func getData(ch chan string) {

	for {
		if input, ok := <-ch; ok {
			fmt.Println(input)
		}
	}
}
func GoroutineTest2() {
	ch := make(chan string)
	go sendData(ch)
	go getData(ch)
	time.Sleep(10 * time.Second)
	fmt.Println("Ending the channel test")
}
