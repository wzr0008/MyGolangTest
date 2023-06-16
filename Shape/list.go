package Shape

import "fmt"

type List []int

func (l *List) Append(val int) {
	*l = append(
		*l,
		val,
	)
}
func (l List) Len() {
	fmt.Println(len(l))
}
