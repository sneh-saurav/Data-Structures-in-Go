package main

import (
	"fmt"
)

type Stack struct {
	elements []int
}

func (s *Stack) Push(i int) {
	s.elements = append(s.elements, i)
}
func (s *Stack) Pop() int {
	l := len(s.elements) - 1
	toRemove := s.elements[l]
	s.elements = s.elements[:l]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)
	myStack.Pop()
	fmt.Println(myStack)
}
