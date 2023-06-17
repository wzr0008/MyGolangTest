package interfaceTest

import (
	"StepbyStep/SortTest"
	"fmt"
)

type Any interface{}

func checkInterface(inter interface{}) {
	switch inter.(type) {
	case int:
		fmt.Println("This is an Integer variable")
	case *SortTest.Person:
		fmt.Println("This is a Person variable")
	case string:
		fmt.Println("This is a string variable ")
	default:
		fmt.Println("We do not know the variable")
	}
}
func TestAny() {
	var v1 Any
	v1 = 5
	checkInterface(v1)
	p := new(SortTest.Person)
	v1 = p
	checkInterface(v1)

}
