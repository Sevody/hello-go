package main

import (
	"fmt"
)

func main() {
	factor:= 2
	s := []int{1,2,3,4,5}
	s2 := Ext(s, factor)
	fmt.Print(s2)
}

func Ext(s []int, factor int) []int {
	ns := make([]int, len(s)*factor)
	copy(ns, s)
	s = ns
	return s
}