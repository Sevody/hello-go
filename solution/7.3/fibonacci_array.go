package main

import "fmt"

func main() {
	var arr [50]int
	arr[0] = 1
	arr[1] = 1
	for i := 2; i < len(arr); i++ {
		arr[i] = arr[i-1] + arr[i-2]
	}
	fmt.Println(arr)
}
