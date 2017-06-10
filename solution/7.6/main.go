package main

import (
	"fmt"
)

func main() {
	buf := []byte{'a','b','c'}
	n:=1
	s1, s2 := buf[:n], buf[n:]
	fmt.Println(s1,s2)
}