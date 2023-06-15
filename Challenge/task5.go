package Challenge

import "fmt"

type Student struct {
	firstName string
	lastName  string
	age       int
}

func StructTesting() {
	s := &Student{
		firstName: "Rui",
		lastName:  "Wang",
		age:       33,
	}
	fmt.Println(s)
}
