package main

import (
	"fmt"
	"strings"
)

func main() {
	asciiOnly := func(r rune) rune {
		if r > 127 {
			return ' '
		}
		return r
	}
	fmt.Print(strings.Map(asciiOnly, "fdsfsd是否打算sd东方闪电fdefwe"))
}
