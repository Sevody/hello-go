package main

import (
	"fmt"
)

func main() {
	ts := "asdfghjk"
	a, b := split(ts, 3)
	fmt.Println(a,"  ",b)
}

func split(s string, i int) (string, string) {
	return s[:i], s[i:]
}
