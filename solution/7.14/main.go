package main

import (
	"fmt"
)

func main() {
	g := "Google"
	rg := reverse(g)
	fmt.Println(rg)
}

func reverse(s string) string {
	ss := []byte(s)
	sl := len(ss)
	for i := 0; i < sl/2; i++ {
		ss[i], ss[sl-i-1] = ss[sl-i-1], ss[i]
	}
	return string(ss)
}
