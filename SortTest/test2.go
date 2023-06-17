package SortTest

import "strings"

type Person struct {
	FirstName string
	LastName  string
}

func (p *Person) String() string {
	return p.LastName + "  " + p.FirstName
}

type Persons []*Person

func (p Persons) Len() int {
	return len(p)
}
func (p Persons) Swap(a, b int) {
	p[a], p[b] = p[b], p[a]
}

func (p Persons) Less(a, b int) bool {
	if strings.Compare(p[a].LastName, p[b].LastName) > 0 {
		return true
	}
	if strings.Compare(p[a].FirstName, p[b].FirstName) > 0 {
		return true
	}
	return false
}
