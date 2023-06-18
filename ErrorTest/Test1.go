package ErrorTest

import (
	"errors"
	"fmt"
)

func ErrorTest1() {
	err := errors.New("Hello world")
	fmt.Println(err.Error())
}
func ErrorTest2() {
	fmt.Println("Starting the project")
	defer func() {
		if e := recover(); e != nil {
			fmt.Println(e)
		}
		fmt.Println("This is the end of the programe")
	}()
	panic("Panic started: stopping the program")

}
func badCall() {
	a, b := 10, 0
	n := a / b
	fmt.Println(n)
}
func ErrorTest3() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Println("Printing the panic message ", e)
		}
	}()
	badCall()
	fmt.Println("After the call")
}
