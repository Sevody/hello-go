package main

import "strconv"

type Stack2 struct {
	index int
	stack Stack
}

func (s *Stack2) Push(a int) {
	s.index += 1
	s.stack.Push(a)
}

func (s *Stack2) Pop() int {
	s.index -= 1
	return s.stack.Pop()
}

func (s *Stack2) String() string {
	return s.stack.String() + " with " + strconv.Itoa(s.index) + "number"
}
