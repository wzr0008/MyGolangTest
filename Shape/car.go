package Shape

import "fmt"

type Engine struct {
}

func (e Engine) Start() {
	fmt.Println("Engine is started")
}
func (e Engine) Stop() {
	fmt.Println("Engine is stopped")
}
func (c Car) Start() {
	fmt.Println("Car is started")
}

type Car struct {
	Engine
}
