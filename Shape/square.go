package Shape

type square struct {
	length int "the length of a square"
	width  int "the width of a square"
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

func NewSquare(length, width int) *square {
	s := new(square)
	s.length = length
	s.width = width
	return s
}
