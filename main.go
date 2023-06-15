package main

import (
	"StepbyStep/Challenge"
	"fmt"
	"strconv"
)

func main() {
	n := 255
	for n > 0 {
		temp := n % 16
		n /= 16
		fmt.Print(strconv.Itoa(temp) + "  ")
	}
	fmt.Println(Challenge.ToFahrenheit(100))
	Challenge.StrTest()
}
