package main

import (
	"fmt"
)

func main() {
	sl := make([]byte,0,10)
	b := []byte{'a', 'b', 'c', 'b', 'c', 'b', 'c', 'b', 'c', 'b', 'c', 'b', 'c'}
	s2 := Append(sl, b)
	fmt.Println(string(s2))
}

func Append(slice []byte, data []byte) []byte {
	lr := len(slice)
	if cap(slice)-lr <= len(data) {
		slice2 := make([]byte, lr+len(data))
		copy(slice2, slice)
		slice = slice2
	}
	copy(slice[lr:], data)
	return slice
}
