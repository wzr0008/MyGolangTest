package Concurrency

import (
	"fmt"
	"time"
)

func longWait() {
	fmt.Println("Beginning the long wait")
	time.Sleep(5 * time.Second)
	fmt.Println("Ending the long wait")
}
func shortwait() {
	fmt.Println("Beginning the short wait")
	time.Sleep(3 * time.Second)
	fmt.Println("Ending the shrot wait")
}
func GoroutineTest() {
	fmt.Println("In main")
	go longWait()
	go shortwait()
	fmt.Println("About to sleep in main")
	time.Sleep(10 * time.Second)
	fmt.Println("At the end of main")
}
