package SortTest

type MySort []int

func (m MySort) Len() int {
	return len(m)
}
func (m MySort) Less(a, b int) bool {
	return (m)[a] < (m)[b]
}
func (m MySort) Swap(a, b int) {
	(m)[a], (m)[b] = (m)[b], (m)[a]
}
