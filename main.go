package main

import (
	"StepbyStep/Concurrency"
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
	Concurrency.GoroutineTest4()
}
