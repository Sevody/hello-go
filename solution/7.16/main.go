package main

import (
	"fmt"
)

func main() {
	s := []int{3, 5, 1, 5, 3, 8, 4, 6}
	ss := bubble(s)
	fmt.Println(ss)
}

func bubble(s []int) []int {
	for i, _ := range s {
		for j := 0; j < len(s)-i-1; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
}
