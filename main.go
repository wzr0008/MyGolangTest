package main

import (
	"StepbyStep/SortTest"
	"StepbyStep/interfaceTest"
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
	p1 := new(SortTest.Person)
	p1.FirstName = "Rui"
	p1.LastName = "Wang"
	p2 := new(SortTest.Person)
	p2.FirstName = "Yan"
	p2.LastName = "Liu"
	p3 := new(SortTest.Person)
	p3.FirstName = "Ana"
	p3.LastName = "Wang"
	persons := SortTest.Persons{p1, p2, p3}
	sort.Sort(persons)
	fmt.Println(persons)
	interfaceTest.TestAny()
}
