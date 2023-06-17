package main

import (
	"StepbyStep/interfaceTest"
	"fmt"
	"time"
)

type myTime struct {
	t time.Time
}

func (t *myTime) ShowTime() {
	now := time.Now()
	fmt.Println(now.Month())
	fmt.Println(now.Day())
}
func main() {
	var s1 interfaceTest.MyStack
	s1.Push("Brown")
	s1.Push(3.14)
	s1.Push(100)
	for {
		item, err := s1.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
