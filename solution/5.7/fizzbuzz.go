package main

import (
	"fmt"
)

const (
	Fizz     = 3
	Buzz     = 5
	FizzBuzz = 15
)

func main() {
	for i := 1; i <= 100; i++ {
		switch {
		case i%FizzBuzz == 0:
			fmt.Println("FizzBuzz")
		case i%Fizz == 0:
			fmt.Println("Fizz")
		case i%Buzz == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
