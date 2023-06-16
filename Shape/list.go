package Shape

type List []int

func (l *List) Append(val int) {
	*l = append(
		*l,
		val,
	)
}
func (l List) Len() int {
	return len(l)
}

type Lener interface {
	Len() int
}
type Appender interface {
	Append(int)
}

func LongEnough(l Lener) bool {
	return l.Len() > 5
}
func CountInfo(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}
