package main

import "fmt"

type T struct {
	f float32
	int
	string
}

func main() {
	t := T{1.0, 2, "a"}
	fmt.Println(t)
}
