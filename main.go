package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

func main() {
	n := 255
	for n > 0 {
		temp := n % 16
		n /= 16
		fmt.Print(strconv.Itoa(temp) + "  ")
	}
	time := 0
	for time <= 10 {
		fmt.Println(rand.Intn(8))
		time++
	}
}
