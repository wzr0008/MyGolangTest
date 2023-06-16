package main

import (
	"StepbyStep/Shape"
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
	var sharper Shape.Shaper
	c := new(Shape.Circle)
	c.Radius = 3
	sharper = c
	switch t := sharper.(type) {
	case *Shape.Circle:
		fmt.Println("This is the circle")
		fmt.Println(t)
	case *Shape.Rectangle:
		fmt.Println("This is the rectangle")
		fmt.Println(t)
	}
}
