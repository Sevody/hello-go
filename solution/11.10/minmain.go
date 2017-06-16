package main

import (
	"./min"
	"fmt"
)

func ints() {
	data := []int{74, 59, 238, -784, 9845, 959, 905, 0, 0, 42, 7586, -5467984, 7586}
	a := min.IntArray(data)
	m := min.Min(a)
	fmt.Println("The minimum of the array is : %v\n",m)
}

func strings() {
	data:=[] string{"ddd","eee","bbb","ccc","aaa"}
	a:=min.StringArray(data)
	m:=min.Min(a)
	fmt.Println("The minimum if the array is: %v\n", m)
}

func main() {
	ints()
	strings()
}