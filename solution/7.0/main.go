package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fs := filter(s, isEven)
	fmt.Print(fs)
}

func filter(s []int, mapping func(int) bool) []int {
	var ns []int
	for _, v := range s {
		if mapping(v) {
			ns = append(ns, v)
		}
	}
	return ns
}

func isEven(n int) bool {
	return n%2 == 0
}
