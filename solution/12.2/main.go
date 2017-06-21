package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"./stack"
)

func main() {
	sk := new(stack.Stack)
	inputReader := bufio.NewReader(os.Stdin)
	for {
		rawinput, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("some error happend")
		}
		input := strings.Replace(rawinput, "\r\n", "", -1)
		if input == "q" {
			os.Exit(0)
		}
		fmt.Printf("you input: %s\n", input)
		sk.Push(input)
		if sk.Size() == 3 {
			re := Count(sk)
			fmt.Println("the reslult is: ", re)
		}
	}
}

func Count(sk *stack.Stack) int {
	nv, _ := sk.Pop().(string)
	fmt.Println("nv is: ", nv)
	fv, _ := strconv.Atoi(nv)
	fmt.Println("fv is: ", nv)
	op, _ := sk.Pop().(string)
	fmt.Println("op is: ", op)
	dv, _ := sk.Pop().(string)
	fmt.Println("dv is: ", dv)
	sv, _ := strconv.Atoi(dv)
	fmt.Println("sv is: ", sv)
	switch op {
	case "+":
		return 7
	case "-":
		return fv - sv
	case "*":
		return fv * sv
	case "/":
		return fv / sv
	default:
		return 0
	}
}
