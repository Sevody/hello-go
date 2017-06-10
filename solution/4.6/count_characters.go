package main

import (
	"fmt"
	"unicode/utf8"
)

const st1 = "asSASA ddd dsjkdsjs dk"
const st2 = "asSASA ddd dsjkdsjsこん dk"

func main() {
	countByte(st1)
	countChar(st2)
}

func countByte(s string) {
	fmt.Println("string has %v len", len(s))
}

func countChar(s string) {
	fmt.Println("string has %v char", utf8.RuneCountInString(s))
}
