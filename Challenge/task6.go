package Challenge

import "fmt"

type Employee struct {
	salary float64
}

func (e *Employee) SetSalary(salary float64) {
	e.salary = salary
}
func (e *Employee) GetSalary() {
	fmt.Println(e.salary)
}
func (e *Employee) RaiseSalary(pct float64) {
	e.salary = e.salary * (1 + pct/100)
	return
}
