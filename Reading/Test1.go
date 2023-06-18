package Reading

import (
	"fmt"
	"strconv"
)

func ReadingTest1() {
	var name string
	var age int
	fmt.Scanln(&name)
	fmt.Scanln(&age)
	fmt.Println("The name is " + name)
	fmt.Println("The age is " + strconv.Itoa(age))
}
