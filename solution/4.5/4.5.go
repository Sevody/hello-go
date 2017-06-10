package main

import "fmt"

type ST string

func main() {
	var a, b ST = "aa", "bb"
	c := a + b
	fmt.Printf("c has value: %v", c)
}
