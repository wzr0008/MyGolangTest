package Shape

import "strconv"

type Square struct {
	length int
	width  int
}

type InnerS struct {
	In1 int
	In2 int
}
type Outers struct {
	B int
	C float32
	InnerS
}

func NewSquare(length, width int) *Square {
	s := new(Square)
	s.length = length
	s.width = width
	return s
}
func (s *Square) String() string {
	return "The length is " + strconv.Itoa(s.length) + ".\nAnd the width is " + strconv.Itoa(s.width)
}
