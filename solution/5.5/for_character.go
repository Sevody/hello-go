package main

import "fmt"

func main() {
	for i := 1; i <= 25; i++ {
		for j := i; j > 0; j-- {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}
	g := "G"
	for i := 1; i <= 25; i++ {
		println(g)
		g += "G"
	}
}
