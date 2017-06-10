package main

import (
	"fmt"
)

func main() {
	pos := 10
	re, p := fibonacci(pos)
	fmt.Printf("fibonacci(%d) is %d", p, re)
}

func fibonacci(n int) (result, pos int) {
	if n <= 1 {
		result = 1
	} else {
		r1, _ := fibonacci(n - 1)
		r2, _ := fibonacci(n - 2)
		result = r1 + r2
	}
	pos = n
	return
}
