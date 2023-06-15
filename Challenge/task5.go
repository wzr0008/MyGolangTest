package Challenge

import (
	"fmt"
	"strings"
)

type Student struct {
	firstName string
	lastName  string
	age       int
}
type Interval struct {
	start int
	end   int
}

func (s *Student) ToUpper() {
	s.firstName = strings.ToUpper(s.firstName)
	s.lastName = strings.ToUpper(s.lastName)
}
func SimpleToUpper(stu Student) {
	stu.lastName = strings.ToUpper(stu.lastName)
	stu.firstName = strings.ToUpper(stu.firstName)
}
func StructTesting() {
	s := &Student{
		firstName: "Rui",
		lastName:  "Wang",
		age:       33,
	}
	SimpleToUpper(*s)
	fmt.Println("we just do a simple try")
	fmt.Println(s)
	s.ToUpper()
	fmt.Println(s)
}
