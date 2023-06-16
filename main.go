package main

import (
	"StepbyStep/SortTest"
	"fmt"
	"sort"
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
	//var lst0 Shape.List
	//if !Shape.LongEnough(lst0) {
	//	fmt.Println("Failed")
	//}
	//lst := new(Shape.List)
	//Shape.CountInfo(lst, 1, 100)
	//if Shape.LongEnough(lst) {
	//	fmt.Println("Done")
	//}
	m := SortTest.MySort{3, 4, 1, 2, 0}
	sort.Sort(m)
	fmt.Println(m)
}
