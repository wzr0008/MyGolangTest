package interfaceTest

import "errors"

type MyStack []interface{}

func (s MyStack) Len() int {
	return len(s)
}
func (s MyStack) IsEmpty() bool {
	return len(s) == 0
}
func (s *MyStack) Push(e interface{}) {
	*s = append(*s, e)
}
func (s MyStack) Top() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is Empty")
	}
	return s[len(s)-1], nil
}
func (s *MyStack) Pop() (interface{}, error) {
	if s.IsEmpty() {
		return nil, errors.New("Stack is Empty")
	}
	val := (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return val, nil
}
