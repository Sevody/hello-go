package main

import (
	"fmt"
)


type Integer int

func (p Integer) get() int {
	return int(p)
}

func main() {
	var i Integer
	i = 22
	fmt.Println(i.get())
}