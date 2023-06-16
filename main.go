package main

import (
	"StepbyStep/Challenge"
	"StepbyStep/Shape"
	"fmt"
	"strconv"
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
	n := 255
	for n > 0 {
		temp := n % 16
		n /= 16
		fmt.Print(strconv.Itoa(temp) + "  ")
	}
	squ := Shape.NewSquare(3, 4)
	fmt.Println(squ)
	ele := new(Shape.Outers)
	ele.B = 3
	ele.C = 4
	ele.InnerS.In1 = 3
	ele.InnerS.In2 = 4

	fmt.Println(ele)
	t := new(myTime)
	t.ShowTime()
	e := new(Challenge.Employee)
	e.SetSalary(100)
	e.RaiseSalary(30)
	e.GetSalary()
}
