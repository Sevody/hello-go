package main

import (
	"fmt"
)

const Len = 20

func main() {
	slice := make([]int, Len)
	getFibonacci(slice)
	fmt.Println(slice)
}
func getFibonacci(s []int) {
	s[0] = 1
	s[1] = 1
	for i := 2; i < len(s); i++ {
		s[i] = s[i-1] + s[i-2]
	}
}
