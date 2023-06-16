package Shape

import (
	"fmt"
	"math"
)

type Rectangle struct {
	Side float64
}
type Circle struct {
	Radius float64
}
type Shaper interface {
	Area()
}

func (r *Rectangle) Area() {
	fmt.Println(r.Side * r.Side)
}
func (c *Circle) Area() {
	fmt.Println(c.Radius * c.Radius * math.Pi)
}
