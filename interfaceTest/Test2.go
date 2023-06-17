package interfaceTest

import "fmt"

type AbsInterface interface {
	Abs(int) int
}
type SqrInterface interface {
	Sqr(int) int
}
type IDuck interface {
	Quack()
	Walk()
}
type Bird struct {
}

func (b Bird) Quack() {
	fmt.Println("Bird is quacking")
}
func (b Bird) Walk() {
	fmt.Println("Bird is Walking")
}
func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Walk()
		duck.Quack()
	}
}
