package main

import "strconv"
import "fmt"

type Stack [4]int

func (s *Stack) Push(a int) {
	for i, v := range s {
		if v == 0 {
			s[i] = a
			break
		}
	}
}

func (s *Stack) Pop() int {
	var re int
	for i := len(s) - 1; i >= 0; i-- {
		if re = s[i]; re != 0 {
			s[i] = 0
			return re
		}
	}
	return 0
}

func (s *Stack) String() string {
	str := ""
	for ix, v := range s {
		str += "[" + strconv.Itoa(ix) + ":" + strconv.Itoa(v) + "]"
	}
	return str
}

func main() {
	s := new(Stack2)
	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Pop()
	s.Push(4)
	fmt.Println(s)
}
