package main

import (
	"fmt"
)

func main() {
	t := []int{1, 2, 3, 4, 5, 6}
	f := func(i int) int {
		return i * 10
	}
	mt := mapf(f, t...)
	fmt.Println(mt)
}

func mapf(f func(int) int, t ...int) []int {
	ret := make([]int, len(t))
	for i, v := range t {
		ret[i] = f(v)
	}
	return ret
}
