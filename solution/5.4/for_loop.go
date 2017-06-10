package main

import "fmt"

func main() {
	for i := 0; i < 15; i++ {
		fmt.Printf("%d in for loop\n", i)
	}
	i := 0
START:
	if i < 15 {
		fmt.Printf("%d in goto\n", i)
		i++
		goto START
	}
}
