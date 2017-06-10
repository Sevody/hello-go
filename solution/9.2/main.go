package main

import (
	"unsafe"
	"fmt"
)

func main() {
	var i int = 10
	si := unsafe.Sizeof(i)
	fmt.Println(si)
}