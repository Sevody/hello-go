package main

import (
	"fmt"
	"os"
	"bufio"
	"strings"
)
var nchar, nword, nline int
func main() {
	nchar, nword, nline = 0, 0, 0
	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("please enter some word, finish with 'S'")
	for {
		input, err := inputReader.ReadString('\n')
		if err != nil {
			fmt.Println("errors")
			os.Exit(1)
		}

		if input == "S\r\n" {
			fmt.Println("the N of characters: ", nchar)
			fmt.Println("the N of word: ", nword)
			fmt.Println("the N of line: ", nline)
			os.Exit(0)
		}
		Counter(input)
	}
}

func Counter(input string) {
	nchar += len(input) - 2
	nword += len(strings.Fields(input))
	nline++
}