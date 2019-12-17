package main

import (
	"bytes"
	"fmt"
)

type stack []int

func (s stack) Push(v int) stack {
	return append(s, v)
}
func (s stack) Pop() (stack, int) {
	l := len(s)
	return s[:l-1], s[l-1]
}
func (s stack) String() string {
	var result bytes.Buffer
	for index := range s {
		c, p := s.Pop()
		result.WriteString("[" + string(index) + string(p) + "]")

		s = c
	}
	return fmt.Sprintf(result.String())
}
func main() {
	s := make(stack, 0)
	s = s.Push(1)
	s = s.Push(2)
	s = s.Push(3)

	//s, p := s.Pop()
	//fmt.Println(p)
	fmt.Println(s)

}
